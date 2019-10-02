package main

import (
	"fmt"
	"routes"

	"github.com/kataras/iris"

	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// db := session.DB("standbymik")
	// controllers := mvc.New(app)
	// controllers.Register(db)
	// controllers.Handle(&models.MyController{})

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

	app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))
}
