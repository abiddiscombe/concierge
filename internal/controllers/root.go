package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RootGetResponse struct {
	Title string                 `json:"title"`
	Links []RootGetResponseLinks `json:"links"`
}

type RootGetResponseLinks struct {
	Href    string `json:"href"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

func RootGet(c echo.Context) error {
	return c.JSON(http.StatusOK, RootGetResponse{
		Title: "[Concierge] Root",
		Links: []RootGetResponseLinks{
			{
				Href:    "/",
				Title:   "Root (Self)",
				Summary: "[GET] Returns information about this API.",
			},
			{
				Href:    "/link",
				Title:   "Link & Alias Management",
				Summary: "[GET, POST] Lookup or create new aliases.",
			},
			{
				Href:    "/to/:alias",
				Title:   "Link Activation & Redirection",
				Summary: "[GET] Accepts a valid alias and redirects to target URL",
			},
		},
	})
}
