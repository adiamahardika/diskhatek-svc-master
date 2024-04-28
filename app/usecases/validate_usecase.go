package usecases

import (
	"context"
	"fmt"
	"math"
	"svc-master/app/helpers"
	"svc-master/app/models"
	customErrors "svc-master/pkg/customerrors"
)

type validateUsecase usecase

type ValidateUsecase interface {
	IsValidCreateStockTransfers(ctx context.Context, request map[string]any) error
	IsValidWarehouse(ctx context.Context, request map[string]any) error
	IsValidUser(ctx context.Context, request map[string]any) error
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
	if srcWh.Status == "inactive" {
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
	if destWh.Status == "inactive" {
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

func (u *validateUsecase) IsValidWarehouse(ctx context.Context, request map[string]any) error {

	warehouse, err := u.Options.Repository.Warehouse.GetDetailWarehouse(ctx, helpers.StrToInt(fmt.Sprint(request["id"])))
	if err != nil {
		return customErrors.NewInternalServiceError(err.Error())
	}
	if warehouse.WarehouseId == 0 {
		return customErrors.NewBadRequestErrorf("Warehouse id %s not found", fmt.Sprint(request["id"]))
	}

	shop, err := u.Options.Repository.Shop.GetDetailShop(ctx, helpers.StrToInt(fmt.Sprint(request["shop_id"])))
	if err != nil {
		return customErrors.NewInternalServiceError(err.Error())
	}
	if shop.ShopId == 0 {
		return customErrors.NewBadRequestErrorf("Shop id %s not found", fmt.Sprint(request["shop_id"]))
	}

	return nil
}

func (u *validateUsecase) IsValidUser(ctx context.Context, request map[string]any) error {

	_, emailPagination, err := u.Options.Repository.User.GetUser(ctx, models.GetUserRequest{
		Email: fmt.Sprint(request["email"]),
		StandardGetRequest: models.StandardGetRequest{
			Page:  1,
			Limit: math.MaxInt32,
		},
	})
	if err != nil {
		return customErrors.NewInternalServiceError(err.Error())
	}
	if emailPagination.Total > 0 {
		return customErrors.NewBadRequestErrorf("Duplicate email! %s already registered.", fmt.Sprint(request["email"]))
	}

	_, phonePagination, err := u.Options.Repository.User.GetUser(ctx, models.GetUserRequest{
		Phone: fmt.Sprint(request["phone"]),
		StandardGetRequest: models.StandardGetRequest{
			Page:  1,
			Limit: math.MaxInt32,
		},
	})
	if err != nil {
		return customErrors.NewInternalServiceError(err.Error())
	}
	if phonePagination.Total > 0 {
		return customErrors.NewBadRequestErrorf("Duplicate phone! %s already registered.", fmt.Sprint(request["phone"]))
	}

	return nil
}
