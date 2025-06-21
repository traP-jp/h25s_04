package repository

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db     *sqlx.DB
	client *s3.Client
}

func New(db *sqlx.DB, client *s3.Client) *Repository {
	return &Repository{db: db, client: client}
}
