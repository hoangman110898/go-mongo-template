package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-mongo-template/src/config"
	"go-mongo-template/src/handlers"
	products "go-mongo-template/src/services"
	"gopkg.in/mgo.v2"
)

func ProductRoutes(app *fiber.App, dbSession *mgo.Session, conf *config.Configuration) {
	productService := products.New(dbSession, conf)
	productHandler := handlers.GetProductHandler(productService)

	// -> middleware logs ip time and status
	api := app.Group("/api", logger.New())

	api.Post("/products", productHandler.Create)
	api.Get("/products", productHandler.GetAll)
	api.Put("/products/:id", productHandler.Update)
	api.Get("/products/:id", productHandler.Detail)
	api.Delete("/products/:id", productHandler.Delete)
}
