package mongoclient

import "gopkg.in/mgo.v2"

// MongoController struct
type MongoController struct {
	session *mgo.Session
}

// MongoSession get session mongodb
func MongoSession() (*mgo.Database, *mgo.Session) {
	session, err := mgo.Dial("localhost")
	if nil != err {
		panic(err)
	}
	// defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("standbymik")

	return c, session
}
