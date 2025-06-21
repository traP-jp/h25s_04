package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/traP-jp/h25s_04/server/cmd/server/server"
	"github.com/traP-jp/h25s_04/server/internal/schema"
	"github.com/traP-jp/h25s_04/server/pkg/config"
	"github.com/traP-jp/h25s_04/server/pkg/database"
	"github.com/traP-jp/h25s_04/server/pkg/storage"
)

func main() {
	e := echo.New()

	// middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8082", // Swagger UI
		},
	}))

	// connect to and migrate database
	db, err := database.Setup(config.MySQL())
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	storageCfg, storageOptFn, err := config.ObjectStorage(context.Background())
	if err != nil {
		e.Logger.Fatal(err)
	}

	storageClient := storage.Setup(storageCfg, storageOptFn)

	s := server.Inject(db, storageClient)
	schema.RegisterHandlersWithBaseURL(e, s.Handler, "/api/v1")

	e.Logger.Fatal(e.Start(config.AppAddr()))
}
