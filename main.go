package main

import (
	"github.com/l1huanyu/x1aol1system/handler"
	"github.com/l1huanyu/x1aol1system/start"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	start.Init()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/missions", handler.GetMissions)
	e.Logger.Fatal(e.Start(":8080"))
}
