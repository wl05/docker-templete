package model

import (
	"time"
	"github.com/globalsign/mgo/bson"
)

type User struct {
    BaseModel
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Avatar    string        `json:"avatar"`
	UserName  string        `json:"username"`
	PassWord  string        `json:"password"`
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}
