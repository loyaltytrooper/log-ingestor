package ping

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func PingService(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
