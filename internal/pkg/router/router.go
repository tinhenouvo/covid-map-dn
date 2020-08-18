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
	var divided [][]model.DataCovid
	chunkSize := 1000
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		divided = append(divided, data[i:end])
	}
	for i := range divided {
		_, err := a.db.Model(&divided[i]).Insert()
		if err != nil {
			print(err)
		}
	}
	print("success")
}

func (a *App) configMiddleware() {
	a.echo.Use(middleware.CORS())
	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())
}
