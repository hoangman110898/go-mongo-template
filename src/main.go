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
	router.ProductRoutes(app, dbSession, conf)
	app.Listen(conf.Address)
}
