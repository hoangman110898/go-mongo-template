package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

type Cookie struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path"`
	Domain   string    `json:"domain"`
	MaxAge   int       `json:"max_age"`
	Expires  time.Time `json:"expires"`
	Secure   bool      `json:"secure"`
	HTTPOnly bool      `json:"http_only"`
	SameSite string    `json:"same_site"`
}

func CookieRouter(app *fiber.App) {
	api := app.Group("/cookie", logger.New())

	api.Get("/set", func(c *fiber.Ctx) error {
		cookie := new(fiber.Cookie)
		cookie.Name = "token"
		cookie.Value = "randomValue"
		cookie.Expires = time.Now().Add(24 * time.Minute)

		c.Cookie(cookie)
		return c.SendString("write cookie done!")
	})

	api.Get("/delete", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Expires:  time.Now().Add(-(time.Hour * 2)),
			HTTPOnly: true,
			SameSite: "lax",
		})

		return c.SendString("delete cooke done! :)")
	})

	api.Get("/", func(c *fiber.Ctx) error {
		value := c.Cookies("token")
		return c.SendString("value from cookie is: " + value)
	})
}
