package usecases

import (
	"svc-master/app/repositories"
	"svc-master/pkg/config"
)

type Main struct {
	Shop ShopUsecase
}

type usecase struct {
	Options Options
}

type Options struct {
	Repository *repositories.Main
	Config     *config.Config
}

func Init(opts Options) *Main {
	uscs := &usecase{opts}

	m := &Main{
		Shop: (*shopUsecase)(uscs),
	}

	return m
}
