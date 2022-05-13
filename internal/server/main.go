package server

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Start(port int, debugMode bool) {
	engine := html.New("./web/template", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/asset", "./web/asset")

	route := "/"
	if debugMode {
		route = "/*"
	}

	app.Get(route, func(ctx *fiber.Ctx) error {
		var page *Page
		var found bool

		if debugMode {
			page, found = pagesMap[strings.TrimPrefix(ctx.Path(), "/")]
		} else {
			page, found = pagesMap[ctx.Get("Host")]
		}

		if !found {
			return ctx.Render("index", pages)
		}
		return ctx.Render("page", page)
	})

	app.Listen(fmt.Sprintf(":%d", port))
}
