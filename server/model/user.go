package model

import (
	"server/db"
	"time"
	"github.com/globalsign/mgo/bson"
	"fmt"
)

type UserModel struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
	Username  string        `json:"username" bson:"username"`
	Email     string        `json:"email" bson:"email"`
	Avatar    string        `json:"avatar" bson:"avatar"`
	Password  string        `json:"password" bson:"password"`
}

type UserInfo struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
	Username  string        `json:"username" bson:"username"`
	Email     string        `json:"email" bson:"email"`
	Avatar    string        `json:"avatar" bson:"avatar"`
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return db.DB.C("users").Insert(&u)
}
// Get user by username
func GetUserByName(username string) (UserModel,error){
	var u UserModel
	err := db.DB.C("users").Find(bson.M{
		"username": username,
	}).One(&u)
	fmt.Println("=====",err)
	return u,err
}


// Get user by id
func GetUserById(id bson.ObjectId) (UserInfo,error){
	var u UserInfo
	err := db.DB.C("users").Find(bson.M{
		"_id": id,
	}).One(&u)
	return u,err
}
