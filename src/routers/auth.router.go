package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/jwt/v2"
	"go-mongo-template/src/config"
	"go-mongo-template/src/handlers"
	"go-mongo-template/src/services"
	"gopkg.in/mgo.v2"
)

func AuthRouter(app *fiber.App, dbSession *mgo.Session, conf *config.Configuration) {
	userService := services.NewUser(dbSession, conf)
	userHandler := handlers.GetUserHandler(userService)

	api := app.Group("/auth", logger.New())

	api.Post("/login", userHandler.Login)
	api.Post("/register", userHandler.Create)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    services.GetPrivateKey().Public(),
	}))

	api.Get("/get-name", userHandler.Detail)
}
