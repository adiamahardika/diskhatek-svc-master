package app

import (
	"svc-master/app/controllers"
	"svc-master/app/repositories"
	"svc-master/app/routes"
	"svc-master/app/usecases"
	"svc-master/pkg/config"
	"svc-master/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Main struct {
	cfg        *config.Config
	database   Database
	repo       *repositories.Main
	usecase    *usecases.Main
	controller *controllers.Main
	router     *echo.Echo
}

type Database struct {
	MySQL    *gorm.DB
	Postgres *gorm.DB
}

func New() *Main {
	return new(Main)
}

func (m *Main) Init() (err error) {
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {

		m.cfg = &config.Config{
			ServiceHost:        "localhost",
			ServiceEndpointV:   "V1",
			ServiceEnvironment: "DEVELOPMENT",
			ServicePort:        "3002",
			Database:           config.DatabasePlatform{},
		}

	} else {

		m.cfg = config.NewConfig()

		m.database.Postgres, err = database.GetConnection(m.cfg.Postgres().Read.ToArgs(database.Postgres, database.ReadConn, nil))
		if err != nil {
			return
		}

	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	m.repo = repositories.Init(repositories.Options{
		Config:   m.cfg,
		Postgres: m.database.Postgres,
	})

	m.usecase = usecases.Init(usecases.Options{
		Config:     m.cfg,
		Repository: m.repo,
	})
	m.controller = controllers.Init(controllers.Options{
		Config:   m.cfg,
		UseCases: m.usecase,
	})

	m.router = e

	routes.Router(e, m.controller)
	return nil
}

func (m *Main) Run() (err error) {
	defer m.close()

	m.router.Start(":" + m.cfg.ServicePort)
	return
}

func (m *Main) close() {
	if m.database.MySQL != nil {
		if db, err := m.database.MySQL.DB(); err == nil {
			db.Close()
		}
	}

	if m.database.Postgres != nil {
		if db, err := m.database.Postgres.DB(); err == nil {
			db.Close()
		}
	}
}
