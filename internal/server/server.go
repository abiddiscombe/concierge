package server

import (
	"fmt"

	"github.com/abiddiscombe/concierge/internal/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	server := echo.New()
	server.HideBanner = true
	server.Pre(middleware.RemoveTrailingSlash())

	server.GET("/", controllers.RootGet)

	server.GET("/to", controllers.ToGet)
	server.GET("/to/:alias", controllers.ToGet)

	server.GET("/link", controllers.LinkGet)
	server.POST("/link", controllers.LinkPost)
	server.GET("/link/:alias", controllers.LinkGet)
	server.POST("/link/:alias", controllers.LinkPost)

	fmt.Println("[Concierge] Server Starting.")
	server.Start(":3000")
}
