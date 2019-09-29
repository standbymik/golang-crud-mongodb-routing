package routes

import (
	"fmt"
	"models"
	"mongoclient"
	"time"

	"github.com/kataras/iris/context"
)

// GetAllUser func for return all user
func GetAllUser(ctx context.Context) {

	db, session := mongoclient.MongoSession()
	defer session.Close()

	c := db.C("users")
	u := []models.User{}
	c.Find(nil).All(&u)

	ctx.JSON(context.Map{"result": u})
}

//AddUser insert user
func AddUser(ctx context.Context) {
	db, session := mongoclient.MongoSession()
	defer session.Close()

	name1 := ctx.FormValue("name")
	age := ctx.FormValue("age")

	// update := bson.M{
	// 	"name":         name,
	// 	"age":          age,
	// 	"created_time": time,
	// }

	update := models.User{Name: name1, Age: age, Createdtime: time.Now()}

	fmt.Println(update)

	err := db.C("users").Insert(update)
	if err != nil {
		ctx.JSON(context.Map{"error": err.Error()})
	}

	ctx.JSON(context.Map{"success": true})
}
