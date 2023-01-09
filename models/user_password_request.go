package models

type UserPasswordRequest struct {
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

func (u *UserPasswordRequest) Validate() bool {
	isPasswordEqual := false
	if u.Password == u.ConfirmPassword {
		isPasswordEqual = true
		return isPasswordEqual
	}
	return isPasswordEqual
}
