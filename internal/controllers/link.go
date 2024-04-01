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
	alias := c.QueryParam("alias")

	if alias == "" {
		return c.JSON(http.StatusBadRequest, LinkResponse{
			Title:   "[Concierge] Alias Lookup",
			Message: "Error. An 'alias' value must be provided.",
		})
	}

	url, createdAt, err := database.LinkRead(alias)

	if url == "" || err != nil {
		errorMessage := fmt.Sprintf("Error. The provided 'alias' of '%s' does not exist.", alias)
		return c.JSON(http.StatusNotFound, LinkResponse{
			Title:   "[Concierge] Alias Lookup",
			Message: errorMessage,
		})
	}

	return c.JSON(http.StatusOK, LinkResponse{
		Title:   "[Concierge] Alias Lookup",
		Message: "Success. Returned information for the alias entry.",
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
	alias := c.QueryParam("alias")

	if url == "" || alias == "" {
		return c.JSON(http.StatusBadRequest, LinkResponse{
			Title:   "[Concierge] Alias Creation",
			Message: "Error. Both 'url' and 'alias' values must be provided.",
		})
	}

	PROTOCOLS := []string{"http://", "HTTP://", "https://", "HTTPS://", "ftp://", "FTP://"}

	for _, value := range PROTOCOLS {
		index := strings.Contains(url, value)
		if index == true {
			return c.JSON(http.StatusBadRequest, LinkResponse{
				Title:   "[Concierge] Alias Creation",
				Message: "Error. The 'url' must not include a protocol (e.g. 'https://').",
			})
		}
	}

	firstCharacter := url[0:1]

	if firstCharacter == "/" {
		return c.JSON(http.StatusBadRequest, LinkResponse{
			Title:   "[Concierge] Alias Creation",
			Message: "Error. The 'url' must start with a fully-qualified domain name.",
		})
	}

	fmt.Println(firstCharacter)

	_, createdAt, err := database.LinkWrite(url, alias)

	if err != nil {
		return c.JSON(http.StatusBadRequest, LinkResponse{
			Title:   "[Concierge] Alias Creation",
			Message: "Error. The provided 'alias' already exists.",
		})
	}

	return c.JSON(http.StatusCreated, LinkResponse{
		Title:   "[Concierge] Alias Creation",
		Message: "Success. A new alias entry has been created.",
		Metadata: &LinkResponseMetadata{
			URL:       fmt.Sprintf("https://%s", url),
			Link:      fmt.Sprintf("https://%s/to/%s", c.Request().Host, alias),
			Alias:     alias,
			CreatedAt: createdAt,
		},
	})
}
