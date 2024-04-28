package routes

import (
	"svc-master/app/controllers"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, controller *controllers.Main) {

	v1 := e.Group("/v1")
	{
		shop := v1.Group("/shop")
		{
			shop.GET("", controller.Shop.Get)
		}
	}
}
