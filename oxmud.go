package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/oxmud/config"
	"github.com/sparkymat/oxmud/database"
	"github.com/sparkymat/oxmud/dbx"
	"github.com/sparkymat/oxmud/internal/route"
)

func main() {
	e := echo.New()

	appConfig, err := config.New()
	if err != nil {
		panic(err)
	}

	dbDriver, err := database.New(appConfig.DatabaseURL())
	if err != nil {
		panic(err)
	}

	if err = dbDriver.AutoMigrate(); err != nil {
		panic(err)
	}

	db := dbx.New(dbDriver.DB())

	route.Setup(e, appConfig, db)

	e.Logger.Panic(e.Start(":8080"))
}
