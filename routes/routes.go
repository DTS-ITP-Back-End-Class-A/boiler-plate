package routes

import (
	"fmt"
	"testing/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes() {
	app := fiber.New()

	app.Get("/siswa/detail/:studentID", , handler.HandlerStudentDetail)

	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {

		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg)
	})

	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg)
	})

	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg)
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg)
	})

	app.Listen(":3000")
}
