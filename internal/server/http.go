package server

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/betterde/ects/api/routes"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/journal"
	"github.com/betterde/ects/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

type HttpServer struct {
	Engine *fiber.App
}

var Instance *HttpServer

func InitHttpServer(name, version string) {
	Instance = &HttpServer{
		Engine: fiber.New(fiber.Config{
			AppName:       name,
			ServerHeader:  fmt.Sprintf("%s %s", name, version),
			CaseSensitive: true,
			// Override default error handler
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				// Status code defaults to 500
				code := fiber.StatusInternalServerError

				// Retrieve the custom status code if it's a fiber.*Error
				var e *fiber.Error
				if errors.As(err, &e) {
					code = e.Code
				}

				if err != nil {
					if code >= fiber.StatusInternalServerError {
						journal.Logger.Errorw("Analysis server runtime error:", zap.Error(err))
					}

					// In case the SendFile fails
					return ctx.Status(code).JSON(response.Send(code, err.Error(), err))
				}

				return nil
			},
		}),
	}

	routes.RegisterRoutes(Instance.Engine)
}

func (s *HttpServer) Run(verbose bool) {
	Instance.Engine.Use(cors.New())

	if verbose {
		Instance.Engine.Use(logger.New())
	}
	Instance.Engine.Use(pprof.New())
	Instance.Engine.Use(recover.New())
	Instance.Engine.Use(requestid.New())

	go func() {
		if config.Conf.HTTP.TLSKey != "" && config.Conf.HTTP.TLSCert != "" {
			cert, err := tls.LoadX509KeyPair(config.Conf.HTTP.TLSCert, config.Conf.HTTP.TLSKey)
			if err != nil {
				journal.Logger.Panicw("Failed to start orbit server:", err)
			}

			// Create custom listener
			ln, err := tls.Listen("tcp", config.Conf.HTTP.Listen, &tls.Config{Certificates: []tls.Certificate{cert}})
			if err != nil {
				journal.Logger.Panicw("Failed to start orbit server:", err)
			}

			err = Instance.Engine.Listener(ln)
			if err != nil {
				journal.Logger.Panicw("Failed to start orbit server:", err)
			}
		} else {
			err := Instance.Engine.Listen(config.Conf.HTTP.Listen)
			if err != nil {
				journal.Logger.Panicw("Failed to start orbit server:", err)
			}
		}
	}()
}
