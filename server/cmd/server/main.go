package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/traP-jp/h25s_04/server/cmd/server/server"
	"github.com/traP-jp/h25s_04/server/pkg/config"
	"github.com/traP-jp/h25s_04/server/pkg/database"
)

func main() {
	e := echo.New()

	// middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// connect to and migrate database
	db, err := database.Setup(config.MySQL())
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	s := server.Inject(db)

	v1API := e.Group("/api/v1")
	s.SetupRoutes(v1API)

	e.Logger.Fatal(e.Start(config.AppAddr()))
}
