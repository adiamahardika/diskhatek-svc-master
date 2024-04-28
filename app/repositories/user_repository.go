package repositories

import (
	"context"
	"svc-master/app/models"
)

type userRepository repository

type UserRepository interface {
	CreateUser(request models.User) (models.User, error)
	GetUser(ctx context.Context, filter models.GetUserRequest) ([]models.User, models.Pagination, error)
}

func (r *userRepository) CreateUser(request models.User) (models.User, error) {

	err := r.Options.Postgres.Create(&request).Error

	return request, err
}

func (r *userRepository) GetUser(ctx context.Context, filter models.GetUserRequest) ([]models.User, models.Pagination, error) {

	var (
		users      []models.User
		pagination models.Pagination
		totalItems int64
	)

	offset := (filter.Page - 1) * filter.Limit

	query := r.Options.Postgres.Table("users").Order("users.name")

	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}

	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}

	if filter.Phone != "" {
		query = query.Where("phone = ?", filter.Phone)
	}

	result := query.Find(&users)

	// count totalItems by filter
	result = result.Count(&totalItems)
	pagination.Total = int(totalItems)

	result = result.WithContext(ctx).Offset(offset).Limit(filter.Limit).Find(&users)
	if result.Error != nil {
		return nil, pagination, result.Error
	}

	// Calculate the total number of pages
	if totalItems%int64(filter.Limit) == 0 {
		pagination.TotalPage = int(totalItems / int64(filter.Limit))
	} else {
		pagination.TotalPage = int(totalItems/int64(filter.Limit)) + 1
	}

	pagination.Page = filter.Page
	pagination.PageSize = filter.Limit

	return users, pagination, nil
}
