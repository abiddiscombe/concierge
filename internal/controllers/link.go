package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/abiddiscombe/concierge/internal/database"
	"github.com/labstack/echo/v4"
)

type LinkResponse struct {
	Title    string                `json:"title"`
	Message  string                `json:"message"`
	Metadata *LinkResponseMetadata `json:"metadata,omitempty"`
}

type LinkResponseMetadata struct {
	URL       string `json:"url"`
	Link      string `json:"link"`
	Alias     string `json:"alias"`
	CreatedAt string `json:"createdAt"`
}

func LinkGet(c echo.Context) error {
	alias := c.Param("alias")

	if alias == "" {
		return echo.NewHTTPError(http.StatusBadRequest, LinkResponse{
			Title:   "[Concierge] Alias Lookup",
			Message: "An 'alias' URL parameter must be provided.",
		})
	}

	url, createdAt, err := database.LinkRead(alias)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, LinkResponse{
			Title:   "[Concierge] Alias Lookup",
			Message: "Internal Server Error",
		})
	}

	if url == "" {
		errorMessage := fmt.Sprintf("The provided alias of '%s' does not exist.", alias)
		return echo.NewHTTPError(http.StatusNotFound, LinkResponse{
			Title:   "[Concierge] Alias Lookup",
			Message: errorMessage,
		})
	}

	return c.JSON(http.StatusOK, LinkResponse{
		Title:   "[Concierge] Alias Lookup",
		Message: "Returned information for the alias entry.",
		Metadata: &LinkResponseMetadata{
			URL:       fmt.Sprintf("https://%s", url),
			Link:      fmt.Sprintf("https://%s/to/%s", c.Request().Host, alias),
			Alias:     alias,
			CreatedAt: createdAt,
		},
	})
}

func LinkPost(c echo.Context) error {
	url := c.QueryParam("url")
	alias := c.Param("alias")

	if url == "" || alias == "" {
		return echo.NewHTTPError(http.StatusBadRequest, LinkResponse{
			Title:   "[Concierge] Alias Creation",
			Message: "Both 'url' and 'alias' parameters must be provided.",
		})
	}

	PROTOCOLS := []string{"http://", "HTTP://", "https://", "HTTPS://", "ftp://", "FTP://"}

	for _, value := range PROTOCOLS {
		index := strings.Contains(url, value)
		if index {
			return echo.NewHTTPError(http.StatusBadRequest, LinkResponse{
				Title:   "[Concierge] Alias Creation",
				Message: "The 'url' must not include a protocol (e.g. 'https://').",
			})
		}
	}

	if url[0:1] == "/" {
		return echo.NewHTTPError(http.StatusBadRequest, LinkResponse{
			Title:   "[Concierge] Alias Creation",
			Message: "The 'url' must start with a fully-qualified domain name.",
		})
	}

	url, createdAt, err := database.LinkWrite(url, alias)

	if err != nil {
		// This approach of determining if an HTTP-500 error
		// has occurred is rather hacky. To be revisted later.
		errorStartingText := err.Error()[0:26]
		if errorStartingText == "ERROR: duplicate key value" {
			return echo.NewHTTPError(http.StatusBadRequest, LinkResponse{
				Title:   "[Concierge] Alias Creation",
				Message: "The specified alias already exists.",
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, LinkResponse{
			Title:   "[Concierge] Alias Creation",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, LinkResponse{
		Title:   "[Concierge] Alias Creation",
		Message: "A new alias entry has been created.",
		Metadata: &LinkResponseMetadata{
			URL:       fmt.Sprintf("https://%s", url),
			Link:      fmt.Sprintf("https://%s/to/%s", c.Request().Host, alias),
			Alias:     alias,
			CreatedAt: createdAt,
		},
	})
}
