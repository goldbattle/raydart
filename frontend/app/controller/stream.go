package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func StreamGET(c echo.Context) error {
	// Get the stream we want to pass
	urlHash := c.Param("stream")
	// Render the view
	return c.Render(http.StatusOK, "stream", urlHash)
}
