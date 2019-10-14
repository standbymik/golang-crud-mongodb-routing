package main

import (
	"fmt"
	"models"
	"routes"

	"github.com/kataras/iris"

	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/get_user", func(ctx context.Context) {
		p := ctx.URLParam("id")
		fmt.Println(p)
	})

	app.Handle("POST", "/check_user2", func(ctx context.Context) {
		p := ctx.PostValue("id")
		fmt.Println(p)
	})

	app.Handle("POST", "/insert_user", routes.AddUser)
	app.Handle("GET", "/get_all_users", routes.GetAllUser)
	app.Handle("GET", "/get_user", routes.GetUser)
	app.Handle("POST", "/update_user", routes.UpdateUser)
	app.Handle("POST", "delete_user", routes.DeleteUser)
	app.Handle("POST", "/test", func(ctx context.Context) {

		user := models.User{}
		ctx.ReadForm(&user)
		ctx.JSON(context.Map{"result": user})
	})
	app.Handle("GET", "/test2", func(ctx context.Context) {
		user := models.User{}
		ctx.ReadQuery(&user)
		ctx.JSON(context.Map{"result": user})
	})

	app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))
}
