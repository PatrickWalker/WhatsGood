package main

import (
	"github.com/PatrickWalker/WhatsGood/app/whatsapp"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"fmt"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	//setupConfig

	wc := whatsapp.NewClient("", "", "testTOKEN123", "")

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(wc)
		wc.Send([]string{"123"}, map[string]string{
			"foo": "bar",
		})
		return c.SendString("Hello, World 👋!")
	})
	//webhook handling

	app.Listen(":3000")
}
