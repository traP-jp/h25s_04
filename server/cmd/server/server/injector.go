package server

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jmoiron/sqlx"

	"github.com/traP-jp/h25s_04/server/internal/handler"
	"github.com/traP-jp/h25s_04/server/internal/repository"
	"github.com/traP-jp/h25s_04/server/internal/schema"
)

type Server struct {
	Handler schema.ServerInterface
}

func Inject(db *sqlx.DB, client *s3.Client) *Server {
	repo := repository.New(db, client)
	h := handler.New(repo)

	return &Server{
		Handler: h,
	}
}
