package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionUser holds the name of the user collection
	CollectionUser = "user"
)

// User model
type User struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string        `json:"email" binding:"required" bson:"email"`
	Password         string        `json:"password" binding:"required" bson:"password"`
}

// LoginReq params
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}