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

type userController controller

type UserController interface {
	Create(ctx echo.Context) error
}

func (c *userController) Create(ctx echo.Context) error {
	var (
		user models.User
		err  error
	)

	req, err := inrequest.Json(ctx.Request())
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	mapReq := req.ToMap()
	schema := validet.NewSchema(
		mapReq,
		map[string]validet.Rule{
			"phone":    validet.String{Required: true},
			"email":    validet.String{Required: true, Email: true},
			"name":     validet.String{Required: true},
			"password": validet.String{Required: true},
		},
		validet.Options{},
	)

	errorBags, err := schema.Validate()
	if err != nil {
		err := customError.NewBadRequestError(err.Error())
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), errorBags.Errors, nil, nil)
	}

	err = req.ToBind(&user)
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	err = c.Options.UseCases.Validate.IsValidUser(ctx.Request().Context(), mapReq)
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}
	user, err = c.Options.UseCases.User.CreateUser(user)
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.StandardResponse(ctx, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, user, nil)
}
