package handlers

import (
	"encoding/json"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-mongo-template/src/models"
	"go-mongo-template/src/services"
)

type UserHandler struct {
	user services.UserService
}

func GetUserHandler(user services.UserService) *UserHandler  {
	return &UserHandler{user}
}

func (u *UserHandler) Create(c *fiber.Ctx) error {
	var user *models.User

	json.Unmarshal([]byte(c.Body()), &user)
	if validateError := user.Validate(); validateError != nil {
		return c.JSON(validateError.Error())
	}

	err := user.Initialize()

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = u.user.FindByEmail(user.Email)
	if  err == nil {
		return c.Status(400).JSON(fiber.Map{
			"message":"this email is used by another account",
		})
	}
	err = u.user.Create(user)

	if err != nil {
		c.Status(400).JSON(
			fiber.Map{
				"message":"we have some problem with create user",
			})
	}

	response, _ := json.Marshal(user)

	return c.Send(response)
}

func (u *UserHandler) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")
	result, err := u.user.Login(email, pass)
	if err != nil {
		return c.Status(400).Send([]byte(err.Error()))
	}
	response, _ := json.Marshal(result)
	return c.Send(response)
}

func (u *UserHandler) Detail(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}