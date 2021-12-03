package handler

import (
	"net/http"

	"github.com/l1huanyu/x1aol1system/constant"
	"github.com/labstack/echo/v4"
)

func GetOneUser(ctx echo.Context) error {
	openID := ctx.Get(constant.OpenID).(string)
	return ctx.String(http.StatusOK, "hello, open id "+openID)
}
