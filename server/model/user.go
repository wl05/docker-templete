package model

import (
	"server/db"
)

type UserModel struct {
	BaseModel
	UserName string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	PassWord string `json:"password"`
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return db.DB.C("users").Insert(u)
}
