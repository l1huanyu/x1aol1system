package main

import (
	"net/http"

	"github.com/l1huanyu/x1aol1system/constant"
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

	v1 := e.Group("api/v1", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			openID := ctx.Request().Header.Get("X-WX-OPENID")
			if openID == "" {
				return ctx.JSON(http.StatusOK, map[string]interface{}{
					"code":    constant.UserNotLogin,
					"message": constant.UserNotLoginMessage,
				})
			}
			ctx.Set(constant.OpenID, openID)
			return next(ctx)
		}
	})
	{
		// 创建新任务
		v1.POST("/missions", handler.CreateOneMission)
		// 获取任务列表
		v1.GET("/missions", handler.GetMissions)
		// 更新已有任务
		v1.PUT("/missions/:id", handler.UpdateOneMission)
		// 删除已有任务
		v1.DELETE("/mission/:id", handler.DeleteOneMission)

		// 完成某任务，创建新任务记录
		v1.POST("/records", handler.CreateOneRecord)
		// 获取任务记录列表
		v1.GET("/records", handler.GetRecords)

		// 创建新商品
		e.POST("/goods", handler.CreateOneGoods)
		// 获取当前用户商品列表
		v1.GET("/goods", handler.GetGoods)
		// 更新已有商品
		e.PUT("/goods", handler.UpdateOneGoods)
		// 删除已有商品
		e.DELETE("/goods", handler.DeleteOneGoods)

		// 购买某商品，创建新订单
		e.POST("/orders", handler.CreateOneOrder)
		// 获取当前用户订单列表
		e.GET("/orders", handler.GetOrders)

		// 获取当前用户个人信息
		e.GET("users", handler.GetOneUser)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
