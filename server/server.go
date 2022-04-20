package server

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/zjyl1994/unilinkd/asset"
	"github.com/zjyl1994/unilinkd/handler"
)

func Run(listenAddr string) error {
	app := fiber.New(fiber.Config{
		ServerHeader: "UniLink",
		Views:        html.NewFileSystem(asset.HttpAssets, ".html"),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := http.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			if err := ctx.SendStatus(code); err != nil {
				return err
			}
			return ctx.Render("tpl", fiber.Map{
				"Title":   http.StatusText(code),
				"Content": err.Error(),
				"Time":    time.Now().Format(time.RFC3339),
			})
		},
	})

	app.Get("/", sendFile("index.html"))
	app.Get("/favicon.ico", sendFile("favicon.ico"))
	app.Get("/+", func(c *fiber.Ctx) error { return handler.CodeHandler(c, c.Params("+")) })
	return app.Listen(listenAddr)
}

func sendFile(path string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		file, err := asset.Assets.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		c.Set("Content-type", http.DetectContentType(data))
		return c.Send(data)
	}
}
