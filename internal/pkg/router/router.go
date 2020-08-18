package router

import (
	"github.com/go-pg/pg/v9"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"crawlerdatacovid/internal/app/model"
	client_request "crawlerdatacovid/internal/pkg/client-request"
	"crawlerdatacovid/internal/pkg/database"
)

type App struct {
	echo *echo.Echo
	db   *pg.DB
}

func NewApp() *App {
	e := echo.New()
	//try to connect postgres db
	db := database.GetConnect()
	return &App{e, db}
}


func (a *App) Run() {
	a.configMiddleware()
	data, _ := client_request.GetData()
	_, err := a.db.Model(model.DataCovid{}).Insert(data)
	if err != nil {
		print (err)
	}
}


func (a *App) configMiddleware() {
	a.echo.Use(middleware.CORS())
	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())
}
