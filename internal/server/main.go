package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Start(port int) {
	engine := html.New("./web/template", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		page, found := pagesMap[ctx.Get("Host")]
		if !found {
			return ctx.Render("index", pages)
		}
		return ctx.Render("page", page)
	})

	app.Static("/asset", "./web/asset")

	app.Listen(fmt.Sprintf(":%d", port))
}
