package usecases

import (
	"context"
	"fmt"
	"svc-master/app/helpers"
	"svc-master/app/models"
	customErrors "svc-master/pkg/customerrors"
)

type validateUsecase usecase

type ValidateUsecase interface {
	IsValidCreateStockTransfers(ctx context.Context, request map[string]any) error
}

func (u *validateUsecase) IsValidCreateStockTransfers(ctx context.Context, request map[string]any) error {

	reqBody := models.StockTransfer{
		ProductId:       helpers.StrToInt(fmt.Sprint(request["product_id"])),
		SrcWarehouseId:  helpers.StrToInt(fmt.Sprint(request["source_warehouse_id"])),
		DestWarehouseId: helpers.StrToInt(fmt.Sprint(request["destination_warehouse_id"])),
		Quantity:        helpers.StrToInt(fmt.Sprint(request["quantity"])),
	}

	srcWh, err := u.Options.Repository.Warehouse.GetDetailWarehouse(ctx, reqBody.SrcWarehouseId)
	if err != nil {
		return customErrors.NewInternalServiceError(err.Error())
	}
	if srcWh.WarehouseId == 0 {
		return customErrors.NewBadRequestErrorf("Warehouse id %d not found", reqBody.SrcWarehouseId)
	}
	if !srcWh.Status {
		return customErrors.NewBadRequestErrorf("Warehouse id %d are inactive", reqBody.SrcWarehouseId)
	}
	if reqBody.Quantity > srcWh.WarehouseId {
		return customErrors.NewBadRequestErrorf("Insufficient stock available at source warehouse id %d", reqBody.SrcWarehouseId)
	}

	destWh, err := u.Options.Repository.Warehouse.GetDetailWarehouse(ctx, reqBody.DestWarehouseId)
	if err != nil {
		return customErrors.NewInternalServiceError(err.Error())
	}
	if destWh.WarehouseId == 0 {
		return customErrors.NewBadRequestErrorf("Warehouse id %d not found", reqBody.DestWarehouseId)
	}
	if !destWh.Status {
		return customErrors.NewBadRequestErrorf("Warehouse id %d are inactive", reqBody.DestWarehouseId)
	}

	product, err := u.Options.Repository.Product.GetDetailProduct(ctx, reqBody.ProductId)
	if err != nil {
		return customErrors.NewInternalServiceError(err.Error())
	}
	if product.ProductId == 0 {
		return customErrors.NewBadRequestErrorf("Product id %d not found", reqBody.ProductId)
	}

	return nil
}
