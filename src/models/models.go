package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// MyController is struct of MongoDB
// type MyController struct {
// 	DB *mgo.Database
// }

// User struct
type User struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Name        string        `json:"name"`
	Content     string        `json:"content"`
	Createdtime time.Time     `json:"created_time" bson:"created_time"`
}
