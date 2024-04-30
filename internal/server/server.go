package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/abiddiscombe/concierge/internal/controllers"
	"github.com/abiddiscombe/concierge/internal/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	logger := log.NewLogger("server")

	server := echo.New()
	server.HidePort = true
	server.HideBanner = true

	server.Pre(middleware.RemoveTrailingSlash())
	server.Use(middleware.Recover())
	server.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogMethod:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			var logLevel slog.Level
			if v.Error == nil {
				logLevel = slog.LevelInfo
			} else if v.Status >= 400 && v.Status <= 499 {
				logLevel = slog.LevelWarn
			} else {
				logLevel = slog.LevelError
			}

			logger.LogAttrs(context.Background(), logLevel, "New HTTP Event",
				slog.Int("status", v.Status),
				slog.String("method", v.Method),
				slog.String("uri", v.URI),
			)

			return nil
		},
	}))

	server.GET("/", controllers.RootGet)

	server.GET("/to", controllers.ToGet)
	server.GET("/to/:alias", controllers.ToGet)

	server.GET("/link", controllers.LinkGet)
	server.POST("/link", controllers.LinkPost)
	server.GET("/link/:alias", controllers.LinkGet)
	server.POST("/link/:alias", controllers.LinkPost)

	if err := server.Start(":3000"); err != http.ErrServerClosed {
		logger.Error("Server closed unexpectedly", "err", err)
	}
}
