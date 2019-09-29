package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User struct
type User struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Age         string        `json:"age" bson:"age"`
	Createdtime time.Time     `json:"created_time,omitempty" bson:"created_time"`
}
