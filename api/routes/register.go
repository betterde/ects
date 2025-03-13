package routes

import (
	"github.com/betterde/ects/api"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/spa"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(response.Success("Success", nil))
	}).Name("Health check")

	// Swagger API specification file router
	app.Get("/swagger/*", filesystem.New(filesystem.Config{
		Root:               api.Serve(),
		Index:              "openapi.yaml",
		NotFoundFile:       "openapi.yaml",
		ContentTypeCharset: "UTF-8",
	})).Name("Swagger JSON Schema")

	// Swagger UI router
	app.Get("/docs/*", swagger.New(swagger.Config{
		URL:          "/swagger/openapi.yaml",
		DeepLinking:  false,
		DocExpansion: "none",
	})).Name("Swagger UI")

	app.Get("*", filesystem.New(filesystem.Config{
		Root:               spa.Serve(),
		Index:              "index.html",
		NotFoundFile:       "index.html",
		ContentTypeCharset: "UTF-8",
	})).Name("SPA static resource")
}
