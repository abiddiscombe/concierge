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
		return echo.NewHTTPError(http.StatusNotFound, ToGetResponse{
			Title:   "[Concierge] Alias Redirection.",
			Message: "An '/alias' value must be provided.",
		})
	}

	url, _, err := database.LinkRead(alias)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ToGetResponse{
			Title:   "[Concierge] Alias Redirection.",
			Message: "Internal Server Error",
		})
	}

	if url == "" {
		return echo.NewHTTPError(http.StatusNotFound, ToGetResponse{
			Title:   "[Concierge] Alias Redirection.",
			Message: "The provided 'alias' is not valid.",
		})
	}

	urlParsed := fmt.Sprintf("https://%s", url)
	return c.Redirect(http.StatusPermanentRedirect, urlParsed)
}
