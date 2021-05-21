package main

import (
	"github.com/gofiber/fiber/v2"
	"go-mongo-template/src/config"
	"go-mongo-template/src/db"
	"go-mongo-template/src/handlers"
	products "go-mongo-template/src/services"
	"gopkg.in/mgo.v2"
)

func main() {
	conf := config.NewConfig()

	dbSession := db.GetInstance(conf)

	dbSession.SetSafe(&mgo.Safe{})

	productService := products.New(dbSession, conf)

	app := fiber.New()
	handlers.ProductRouter(productService, app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Connect success"))
	})
	app.Listen(conf.Address)
}
