package repositories

import (
	"svc-master/app/models"
)

type reservedStockRepository repository

type ReservedStockRepository interface {
	DeleteReservedStock() error
}

func (r *reservedStockRepository) DeleteReservedStock() error {

	var reserved_stocks models.ReservedStock
	err := r.Options.Postgres.Table("reserved_stocks").Where("reservation_expiry_time < NOW()").Delete(&reserved_stocks).Error

	return err
}
