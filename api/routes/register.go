package routes

import (
	"github.com/betterde/ects/api"
	"github.com/betterde/ects/api/handler"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/spa"
	jwtware "github.com/gofiber/contrib/jwt"
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

	app.Post("/api/signin", handler.SignInHandler)
	app.Post("/api/signup", handler.SignUpHandler)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Conf.Auth.Secret)},
	}))

	app.Post("/api/profile", handler.GetProfile)

	app.Get("*", filesystem.New(filesystem.Config{
		Root:               spa.Serve(),
		Index:              "index.html",
		NotFoundFile:       "index.html",
		ContentTypeCharset: "UTF-8",
	})).Name("SPA static resource")
}

func RegisterInitializeRoutes(app *fiber.App) {
	app.Get("/api/system/info", handler.GetSystemInfo)
	app.Post("/api/system/secret", handler.GenSystemSecret)
	app.Post("/api/system/database", handler.SetSystemDatabase)
	app.Post("/api/system/initialization", handler.InitializationSystem)

	// Initialization SPA route
	app.Get("*", filesystem.New(filesystem.Config{
		Root:               spa.Serve(),
		Index:              "initialize.html",
		NotFoundFile:       "initialize.html",
		ContentTypeCharset: "UTF-8",
	})).Name("SPA static resource")
}
