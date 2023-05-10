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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testUser = &entities.User{
	ID:       2,
	Name:     "test_sahan",
	Email:    "test_sahan@hotmail.com",
	Password: "1111",
}
var testUser1 = &entities.User{
	ID:       1,
	Name:     "test_sahan1",
	Email:    "test_sahan1@hotmail.com",
	Password: "1111",
}
var users = []*entities.User{
	testUser, testUser1,
}

func TestSignUser(t *testing.T) {
	testRepository := &mocks.UserRepositoryInterface{}
	testRepository.On("CreateUser", mock.AnythingOfType("*entities.User")).Return(nil)

	newHandler := handler.CreateUserHandler(testRepository)

	requestBody, err := json.Marshal(testUser)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, err = http.NewRequest(http.MethodPost, "/sign_up", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)
	c.Request.Header.Add("Content-Type", "application/json")

	newHandler.SignUser(c)
	assert.Equal(t, http.StatusCreated, w.Code)
	expectedResponseBody := (&entities.User{
		ID:            2,
		Password:      "1111",
		Name:          "test_sahan",
		Email:         "test_sahan@hotmail.com",
		IsEmailActive: false,
	})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponseBody, testUser)

	testRepository.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	testRepository := &mocks.UserRepositoryInterface{}
	testRepository.On("GetUserByID", uint64(1)).Return(testUser1, nil)
	newHandler := handler.CreateUserHandler(testRepository)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var err error
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	c.Request, err = http.NewRequest(http.MethodGet, "/user/:id", nil)
	assert.NoError(t, err)
	c.Request.Header.Add("Content-Type", "application/json")

	newHandler.GetUser(c)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponseBody := &entities.User{
		ID:        1,
		Name:      "test_sahan1",
		Email:     "test_sahan1@hotmail.com",
		Password:  "1111",
		CreatedAt: testUser1.CreatedAt,
		UpdatedAt: testUser1.UpdatedAt,
	}
	assert.NoError(t, err)
	assert.Equal(t, expectedResponseBody, testUser1)

	testRepository.AssertExpectations(t)
}

func TestGetAll(t *testing.T) {
	testRepository := &mocks.UserRepositoryInterface{}
	testRepository.On("GetAllUsers", mock.AnythingOfType("[]*entities.User")).Return(users, nil)
	requestBody, err := json.Marshal(users)
	assert.NoError(t, err)

	newHandler := handler.CreateUserHandler(testRepository)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	assert.NoError(t, err)
	c.Request, err = http.NewRequest(http.MethodGet, "/users", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)
	c.Request.Header.Add("Content-Type", "application/json")
	newHandler.GetAllUsers(c)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponseBody := []*entities.User{
		testUser,
		testUser1,
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponseBody, users)
	testRepository.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	testRepository := &mocks.UserRepositoryInterface{}
	testRepository.On("GetUserByID", uint64(1)).Return(testUser1, nil)
	testRepository.On("DeleteUser", uint64(1)).Return(nil)
	newHandler := handler.CreateUserHandler(testRepository)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var err error
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	c.Request, err = http.NewRequest(http.MethodDelete, "/user/delete/:id", nil)
	assert.NoError(t, err)
	c.Request.Header.Add("Content-Type", "application/json")

	newHandler.DeleteUser(c)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.NoError(t, err)

	testRepository.AssertExpectations(t)
}
