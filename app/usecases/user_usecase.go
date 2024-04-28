package usecases

import (
	"context"
	"math"
	"svc-master/app/helpers"
	"svc-master/app/models"
	customErrors "svc-master/pkg/customerrors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase usecase

type UserUsecase interface {
	CreateUser(request models.User) (models.User, error)
	Login(ctx context.Context, request models.User) (models.LoginResponse, error)
	Authentication(tokenString string) error
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

func (u *userUsecase) Login(ctx context.Context, request models.User) (models.LoginResponse, error) {

	reqBody := models.GetUserRequest{
		Email: request.Email,
		StandardGetRequest: models.StandardGetRequest{
			Page:  1,
			Limit: math.MaxInt32,
		},
	}
	if reqBody.Email == "" {
		reqBody.Phone = request.Phone
	}

	user, _, err := u.Options.Repository.User.GetUser(ctx, reqBody)
	if err != nil {
		return models.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user[0].PasswordHash), []byte(request.Password))
	if err != nil {
		return models.LoginResponse{}, customErrors.NewBadRequestError("Password not match!")
	}

	tokenLifespan := viper.GetInt("TOKEN_LIFESPAN")
	expirationTime := time.Now().Add(time.Minute * time.Duration(tokenLifespan))

	claims := &models.Claims{
		SignatureKey: helpers.GetMD5Hash(user[0].Email, user[0].Phone),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(viper.GetString("TOKEN_SECRET"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return models.LoginResponse{}, err
	}

	return models.LoginResponse{Token: tokenString}, nil
}

func (u *userUsecase) Authentication(tokenString string) error {

	claims := &models.Claims{}
	jwtKey := []byte(viper.GetString("TOKEN_SECRET"))
	token, error := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	validatorError, _ := error.(*jwt.ValidationError)
	if token == nil {
		return customErrors.NewBadRequestError("Please provide token!")
	} else if validatorError != nil && validatorError.Errors == jwt.ValidationErrorExpired {
		return customErrors.NewBadRequestError("Your token expired!")
	} else if error != nil {
		return customErrors.NewBadRequestError("Your token invalid!")
	}

	return nil
}
