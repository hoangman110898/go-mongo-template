package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-mongo-template/src/models"
	"go-mongo-template/src/services"
)

type ProductHandler struct {
	product products.ProductService
}

func ProductRouter(product products.ProductService, app *fiber.App) {
	productHandler := &ProductHandler{product}

	app.Post(BaseRoute+"/products", productHandler.Create)
	app.Get("/products", productHandler.GetAll)
}

func (p *ProductHandler) Create(c *fiber.Ctx) error {
	var product *models.Product
	json.Unmarshal([]byte(c.Body()), &product)

	err := p.product.Create(c.Context(), product)

	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}
	response, _ := json.Marshal(product)
	return c.Send(response)
}

func (p *ProductHandler) GetAll(c *fiber.Ctx) error {
	products, err := p.product.FindAll(c.Context())
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}
	response, _ := json.Marshal(products)
	return c.Send(response)
}
