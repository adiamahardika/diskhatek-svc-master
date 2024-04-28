package controllers

import (
	"svc-master/app/usecases"
	"svc-master/pkg/config"
)

type Main struct {
	Shop ShopController
}

type controller struct {
	Options Options
}

type Options struct {
	Config   *config.Config
	UseCases *usecases.Main
}

func Init(opts Options) *Main {
	ctrlr := &controller{opts}

	m := &Main{
		Shop: (*shopController)(ctrlr),
	}

	return m
}
