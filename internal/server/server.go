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
	server.HideBanner = true

	server.Pre(middleware.RemoveTrailingSlash())
	server.Use(middleware.Recover())
	server.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
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
		logger.Error("server unexpectedly closed")
	}
}
