package controllers

import (
	"fmt"
	"net/http"
	"svc-master/app/constants"
	"svc-master/app/helpers"
	"svc-master/app/models"
	"time"

	customError "svc-master/pkg/customerrors"

	"github.com/ezartsh/inrequest"
	"github.com/ezartsh/validet"
	"github.com/labstack/echo/v4"
)

type stockTransferController controller

type StockTrasferController interface {
	Create(ctx echo.Context) error
}

func (c *stockTransferController) Create(ctx echo.Context) error {
	var (
		stockTransfer models.StockTransfer
		err           error
	)

	req, err := inrequest.Json(ctx.Request())
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	mapReq := req.ToMap()
	if intValue, ok := mapReq["product_id"].(int); ok {
		fmt.Print(intValue)
	} else {
		fmt.Print("error ini")
	}

	schema := validet.NewSchema(
		mapReq,
		map[string]validet.Rule{
			"transfer_date": validet.String{Required: true, Custom: func(v string, path validet.PathKey, look validet.Lookup) error {
				_, err := time.Parse(time.DateOnly, v)
				if err != nil {
					return customError.NewBadRequestError(constants.InvalidDateFormat)
				}

				return nil
			}},
		},
		validet.Options{},
	)

	errorBags, err := schema.Validate()
	if err != nil {
		err := customError.NewBadRequestError(err.Error())
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), errorBags.Errors, nil, nil)
	}

	err = req.ToBind(&stockTransfer)
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	err = c.Options.UseCases.Validate.IsValidCreateStockTransfers(ctx.Request().Context(), mapReq)
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}
	stockTransfer, err = c.Options.UseCases.StockTransfer.CreateStockTransfer(stockTransfer)
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.StandardResponse(ctx, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, stockTransfer, nil)
}
