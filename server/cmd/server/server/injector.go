package server

import (
	"github.com/jmoiron/sqlx"

	"github.com/traP-jp/h25s_04/server/internal/handler"
	"github.com/traP-jp/h25s_04/server/internal/repository"
	"github.com/traP-jp/h25s_04/server/internal/schema"
)

type Server struct {
	Handler schema.ServerInterface
}

func Inject(db *sqlx.DB) *Server {
	repo := repository.New(db)
	h := handler.New(repo)

	return &Server{
		Handler: h,
	}
}
