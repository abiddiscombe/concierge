package server

import (
	"fmt"

	"github.com/abiddiscombe/concierge/internal/controllers"
	"github.com/labstack/echo/v4"
)

func Init() {
	server := echo.New()

	server.HideBanner = true

	server.GET("/", controllers.RootGet)

	server.GET("/to", controllers.ToGet)
	server.GET("/to/:alias", controllers.ToGet)

	server.GET("/link", controllers.LinkGet)
	server.POST("/link", controllers.LinkPost)

	fmt.Println("[Concierge] Server Starting.")

	server.Start(":3000")
}
