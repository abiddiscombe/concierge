package controllers

import (
	"fmt"
	"net/http"

	"github.com/abiddiscombe/concierge/internal/database"
	"github.com/labstack/echo/v4"
)

type ToGetResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func ToGet(c echo.Context) error {
	alias := c.Param("alias")

	if alias == "" {
		return c.JSON(http.StatusNotFound, ToGetResponse{
			Title:   "[Concierge] Alias Missing.",
			Message: "An '/alias' value must be provided.",
		})
	}

	url, _, err := database.LinkRead(alias)

	if url == "" || err != nil {
		return c.JSON(http.StatusNotFound, ToGetResponse{
			Title:   "[Concierge] Alias Invalid.",
			Message: "The '/alias' provided is not valid.",
		})
	}

	urlParsed := fmt.Sprintf("https://%s", url)
	return c.Redirect(http.StatusPermanentRedirect, urlParsed)
}
