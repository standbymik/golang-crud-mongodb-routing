package routes

import (
	"models"
	"mongoclient"

	"github.com/kataras/iris/context"
)

// GetAllUser func for return all user
func GetAllUser(ctx context.Context) {

	db, session := mongoclient.MongoSession()
	defer session.Close()

	c := db.C("forum")
	user := []models.User{}
	c.Find(nil).All(&user)

	ctx.JSON(context.Map{"result": user})
}
