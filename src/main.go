package main

import (
	"github.com/gofiber/fiber/v2"
	"go-mongo-template/src/config"
	"go-mongo-template/src/db"
	"go-mongo-template/src/routers"
	"gopkg.in/mgo.v2"
)

func main() {
	conf := config.NewConfig()

	dbSession := db.GetInstance(conf)

	dbSession.SetSafe(&mgo.Safe{})

	app := fiber.New()
	routers.ProductRoutes(app, dbSession, conf)
	routers.CookieRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"page": "Home Page",
			"author": "Nguyễn Hoàng Mẫn",
		})
	})
	app.Listen(conf.Address)
}
