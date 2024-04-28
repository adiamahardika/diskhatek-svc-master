package helpers

import (
	"svc-master/app/models"

	"github.com/labstack/echo/v4"
)

func ResponseWrapper(c echo.Context, statusCode int, response interface{}) error {
	return c.JSON(statusCode, response)
}

func StandardResponse(c echo.Context, statusCode int, message interface{}, data interface{}, pagination *models.Pagination) error {
	switch {
	case pagination == nil:
		return ResponseWrapper(c, statusCode, models.StandardResponse{
			StatusCode: statusCode,
			Message:    message,
			Data:       data,
		})
	default:
		return ResponseWrapper(c, statusCode, models.StandardResponseWithPaginate{
			StandardResponse: models.StandardResponse{
				StatusCode: statusCode,
				Message:    message,
				Data:       data,
			},
			Pagination: pagination,
		})
	}
}
