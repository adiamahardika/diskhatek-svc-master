package routes

import (
	"svc-master/app/controllers"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, controller *controllers.Main) {

	v1 := e.Group("/v1")
	{

		v1.POST("/register", controller.User.Create)

		shop := v1.Group("/shop")
		{
			shop.GET("", controller.Shop.Get)
		}

		warehouse := v1.Group("/warehouse")
		{
			warehouse.GET("", controller.Warehouse.Get)
			warehouse.PUT("/:id", controller.Warehouse.Update)
		}

		product := v1.Group("/product")
		{
			product.GET("", controller.Product.Get)
		}

		stockTransfer := v1.Group("/stock-transfer")
		{
			stockTransfer.POST("", controller.StockTransfer.Create)
		}
	}
}
