package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/malusev998/dusanmalusev/container"
	"github.com/malusev998/dusanmalusev/handlers"
)

func RegisterRouter(c *container.Container, app *fiber.App) {
	app.Static("/public", "./public", fiber.Static{
		Browse:    false,
		Compress:  true,
		ByteRange: true,
	})


	registerHomeRoutes(c, app.Group(""))
	registerContactRoutes(c, app.Group("/contact"))
}

func registerHomeRoutes(c *container.Container, app fiber.Router) {
	home := handlers.Home{}

	app.Get("/", home.Home)
}


func registerContactRoutes(c *container.Container, router fiber.Router) {
	contact := handlers.Contact{
		ContactService: c.GetContactService(),
	}

	router.Get("/", contact.Index)
	router.Post("/", contact.Message)
}
