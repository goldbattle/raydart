package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func HomeGET(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}
