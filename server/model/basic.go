package model

import (
	"time"
)

type BaseModel struct {
	ID  bson.ObjectId `json:"id" bson:"_id"`
    CreatedAt time.Time     `json:"createdAt"`
    UpdatedAt time.Time     `json:"updatedAt"`
}