package routes

import (
	"svc-master/app/controllers"
	"svc-master/app/usecases"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

func Router(e *echo.Echo, controller *controllers.Main, usecase *usecases.Main) {

	v1 := e.Group("/v1")
	{

		v1.POST("/register", controller.User.Create)
		v1.POST("/login", controller.User.Login)

		shop := v1.Group("/shop")
		{
			shop.Use(controller.User.Authentication())
			shop.GET("", controller.Shop.Get)
		}

		warehouse := v1.Group("/warehouse")
		{
			warehouse.Use(controller.User.Authentication())
			warehouse.GET("", controller.Warehouse.Get)
			warehouse.PUT("/:id", controller.Warehouse.Update)
		}

		product := v1.Group("/product")
		{
			product.Use(controller.User.Authentication())
			product.GET("", controller.Product.Get)
		}

		stockTransfer := v1.Group("/stock-transfer")
		{
			stockTransfer.Use(controller.User.Authentication())
			stockTransfer.POST("", controller.StockTransfer.Create)
		}
	}

	scheduler := cron.New()
	scheduler.AddFunc("*/5 * * * *", usecase.Scheduler.DeleteReservedStock)
	scheduler.Start()
}
