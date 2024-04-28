package controllers

import (
	"net/http"
	"svc-master/app/constants"
	"svc-master/app/helpers"
	"svc-master/app/models"
	customError "svc-master/pkg/customerrors"

	"github.com/ezartsh/inrequest"
	"github.com/ezartsh/validet"
	"github.com/labstack/echo/v4"
)

type warehouseController controller

type WarehouseController interface {
	Get(ctx echo.Context) error
}

func (c *warehouseController) Get(ctx echo.Context) error {
	var (
		request    models.GetWarehouseRequest
		warehouses []models.Warehouse
		pagination models.Pagination
		err        error
	)

	queryReq := inrequest.Query(ctx.Request())
	mapReq := queryReq.ToMap()

	schema := validet.NewSchema(
		mapReq,
		map[string]validet.Rule{},
		validet.Options{},
	)

	errorBags, err := schema.Validate()
	if err != nil {
		err := customError.NewBadRequestError(err.Error())
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), errorBags.Errors, nil, nil)
	}

	err = queryReq.ToBind(&request)
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	warehouses, pagination, err = c.Options.UseCases.Warehouse.GetWarehouse(ctx.Request().Context(), request)
	if err != nil {
		return helpers.StandardResponse(ctx, http.StatusInternalServerError, []string{err.Error()}, nil, nil)
	}
	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, warehouses, &pagination)
}