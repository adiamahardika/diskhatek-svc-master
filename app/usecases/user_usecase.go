package usecases

import (
	"svc-master/app/models"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase usecase

type UserUsecase interface {
	CreateUser(request models.User) (models.User, error)
}

func (u *userUsecase) CreateUser(request models.User) (models.User, error) {

	newPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	request.PasswordHash = string(newPass)

	result, err := u.Options.Repository.User.CreateUser(request)
	result.PasswordHash = ""
	result.Password = ""
	return result, err
}
