package routes

import (
	"fmt"
	"models"
	"mongoclient"
	"strconv"
	"time"

	"github.com/kataras/iris/context"
	"gopkg.in/mgo.v2/bson"
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

	update := models.User{Name: name1, Age: age, Createdtime: time.Now()}

	fmt.Println(update)

	err := db.C("users").Insert(update)
	if err != nil {
		ctx.JSON(context.Map{"error": err.Error()})
	}

	ctx.JSON(context.Map{"success": true})
}

// GetUser return one user
func GetUser(ctx context.Context) {
	name := ctx.URLParam("name")

	user := models.User{}

	db, session := mongoclient.MongoSession()
	defer session.Close()

	err := db.C("users").Find(bson.M{"name": name}).One(&user)

	if err != nil {
		ctx.JSON(context.Map{"success": false})
	} else {
		ctx.JSON(context.Map{"data": user, "success": false})
	}

}

//UpdateUser update by user
func UpdateUser(ctx context.Context) {

	name := ctx.PostValue("name")
	db, session := mongoclient.MongoSession()
	defer session.Close()

	user := models.User{}
	ctx.ReadForm(&user)

	err := db.C("users").Find(bson.M{"name": name}).One(nil)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(context.Map{"result": "not found"})
	} else {
		where := bson.M{"name": user.Name}
		age, _ := strconv.ParseInt(user.Age, 0, 0)
		update := bson.M{"$set": bson.M{
			"age":          age,
			"created_time": time.Now(),
		},
		}
		db.C("users").Update(where, update)
		ctx.JSON(context.Map{"success": true, "result": user})
	}

}

//DeleteUser for delete by user
func DeleteUser(ctx context.Context) {
	user := models.User{}
	ctx.ReadForm(&user)

	db, session := mongoclient.MongoSession()
	defer session.Close()

	where := bson.M{"name": user.Name}
	err := db.C("users").Find(where).One(nil)
	if err != nil {
		ctx.JSON(context.Map{"result": "not found"})
	} else {
		where := bson.M{"name": user.Name}
		db.C("users").Remove(where)
		ctx.JSON(context.Map{"success": true})
	}
}
