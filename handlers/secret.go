package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SecretRoute(c echo.Context) error {
	return c.String(http.StatusOK, "You have accessed a protected route!")
}
