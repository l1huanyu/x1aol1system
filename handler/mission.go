package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMissions(ctx echo.Context) error {
	openID := ctx.Request().Header.Get("X-WX-OPENID")

	return ctx.String(http.StatusOK, "hello, open id "+openID)
}
