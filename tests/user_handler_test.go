package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoapi/entities"
	"todoapi/handler"
	"todoapi/mocks"
	"todoapi/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var testRepository = &mocks.UserRepositoryInterface{}
var newHandler = handler.CreateUserHandler(testRepository)

func getTestUsers() []*entities.User {
	testUser := &entities.User{
		ID:            2,
		Name:          "test_sahan",
		Email:         "test_sahan@hotmail.com",
		IsEmailActive: true,
		Password:      "1111",
	}
	testUser1 := &entities.User{
		ID:       1,
		Name:     "test_nihat",
		Email:    "test_nihat@hotmail.com",
		Password: "4444",
	}
	users := []*entities.User{
		testUser, testUser1,
	}
	return users
}
func getTestPassword() *models.UserPasswordRequest {
	var testPassword = &models.UserPasswordRequest{
		Password:        "5555",
		ConfirmPassword: "5555",
	}
	return testPassword

}

func getTestEmail() *models.UserResetPasswordRequest {
	users := getTestUsers()
	var testUserEmail = &models.UserResetPasswordRequest{
		Email: users[0].Email,
	}
	return testUserEmail
}

func TestSignUser(t *testing.T) {
	testRepository.On("CreateUser", mock.AnythingOfType("*entities.User")).Return(nil)
	users := getTestUsers()

	requestBody, err := json.Marshal(users[0])
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, err = http.NewRequest(http.MethodPost, "/sign_up", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)

	newHandler.SignUser(c)
	assert.Equal(t, http.StatusCreated, w.Code)
	expectedResponseBody := (&entities.User{
		ID:            2,
		Password:      "1111",
		Name:          "test_sahan",
		Email:         "test_sahan@hotmail.com",
		IsEmailActive: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponseBody, users[0])

	testRepository.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	users := getTestUsers()
	testRepository.On("GetUserByID", uint64(1)).Return(users[1], nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var err error
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	c.Request, err = http.NewRequest(http.MethodGet, "/user/:id", nil)
	assert.NoError(t, err)

	newHandler.GetUser(c)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponseBody := &entities.User{
		ID:        1,
		Name:      "test_nihat",
		Email:     "test_nihat@hotmail.com",
		Password:  "4444",
		CreatedAt: users[1].CreatedAt,
		UpdatedAt: users[1].UpdatedAt,
	}
	assert.NoError(t, err)
	assert.Equal(t, expectedResponseBody, users[1])

	testRepository.AssertExpectations(t)
}

func TestGetAll(t *testing.T) {
	users := getTestUsers()
	testRepository.On("GetAllUsers", mock.AnythingOfType("[]*entities.User")).Return(users, nil)
	requestBody, err := json.Marshal(users)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	assert.NoError(t, err)
	c.Request, err = http.NewRequest(http.MethodGet, "/users", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)
	newHandler.GetAllUsers(c)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponseBody := []*entities.User{
		users[0],
		users[1],
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponseBody, users)
	testRepository.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	users := getTestUsers()
	testRepository.On("GetUserByID", uint64(1)).Return(users[1], nil)
	testRepository.On("DeleteUser", uint64(1)).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var err error
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	c.Request, err = http.NewRequest(http.MethodDelete, "/user/delete/:id", nil)
	assert.NoError(t, err)

	newHandler.DeleteUser(c)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.NoError(t, err)

	testRepository.AssertExpectations(t)
}

func TestUpdateUserPassword(t *testing.T) {
	users := getTestUsers()
	testPassword := getTestPassword()
	testRepository.On("GetUserByID", uint64(1)).Return(users[1], nil)
	testRepository.On("UpdateUserPassword", mock.AnythingOfType("*entities.User")).Return(nil)
	requestBody, err := json.Marshal(testPassword)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	c.Request, err = http.NewRequest(http.MethodPatch, "/user/update/:id", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)
	newHandler.UpdateUserPassword(c)
	assert.Equal(t, http.StatusCreated, w.Code)
	err = bcrypt.CompareHashAndPassword([]byte(users[1].Password), []byte(testPassword.Password))
	assert.NoError(t, err)
	/*expectedResponseBody := &entities.User{
		ID:        1,
		Name:      "test_nihat",
		Email:     "test_nihat@hotmail.com",
		Password:  passwordChanging.Password,
		CreatedAt: users[1].CreatedAt,
		UpdatedAt: users[1].UpdatedAt,
	}
	assert.NoError(t, err)
	assert.Equal(t, expectedResponseBody, users[1])*/

	testRepository.AssertExpectations(t)
}

func TestUserResetPassword(t *testing.T) {
	users := getTestUsers()
	testRepository.On("GetUserByEmail", mock.AnythingOfType("string")).Return(users[0], nil)
	testUserEmail := getTestEmail()
	requestBody, err := json.Marshal(testUserEmail)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, err = http.NewRequest(http.MethodPatch, "/resetpassword", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)
	newHandler.UserResetPassword(c)
	assert.Equal(t, http.StatusCreated, w.Code)
	testRepository.AssertExpectations(t)
}

func TestActivateEmail(t *testing.T) {
	users := getTestUsers()
	testRepository.On("GetUserByID", uint64(1)).Return(users[1], nil)
	testRepository.On("UpdateIsEmailActive", mock.AnythingOfType("*entities.User")).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var err error
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	c.Request, err = http.NewRequest(http.MethodPatch, "/activation:id", nil)
	assert.NoError(t, err)
	newHandler.ActivateEmail(c)
	assert.Equal(t, http.StatusCreated, w.Code)
	expectedResponseBody := (&entities.User{
		ID:            1,
		Name:          "test_nihat",
		Email:         "test_nihat@hotmail.com",
		Password:      "4444",
		IsEmailActive: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponseBody, users[1])

	testRepository.AssertExpectations(t)
}
