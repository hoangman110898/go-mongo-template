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

func GetProductHandler(product products.ProductService) *ProductHandler {
	return &ProductHandler{product}
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

func (p *ProductHandler) Update(c *fiber.Ctx) error {
	var product *models.Product

	json.Unmarshal(c.Body(), &product)
	err := p.product.Update(c.Context(), c.Params("id"), product)
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}
	return c.SendString("Update product success")
}

func (p *ProductHandler) Detail(c *fiber.Ctx) error {
	product, err := p.product.FindOneById(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(404).Send([]byte(err.Error()))
	}
	response, _ := json.Marshal(product)

	return c.Send(response)
}

func (p *ProductHandler) Delete(c *fiber.Ctx) error {
	err := p.product.Delete(c.Context(), c.Params("id"))

	if err != nil {
		return c.Status(404).Send([]byte(err.Error()))
	}

	return c.SendString("Delete at id: " + c.Params("id") + " success")
}
