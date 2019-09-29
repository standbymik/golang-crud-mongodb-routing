package mongoclient

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// MongoController struct
type MongoController struct {
	session *mgo.Session
}

// MongoSession get session mongodb
func MongoSession() (*mgo.Database, *mgo.Session) {
	session, err := mgo.Dial("mongodb://root:123456@localhost:27017")
	if nil != err {
		fmt.Println("error")
		panic(err)
	}
	// defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("standbymik")

	return c, session
}
