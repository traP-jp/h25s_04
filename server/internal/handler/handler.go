package handler

import (
	"context"

	"github.com/traPtitech/go-traq"

	"github.com/traP-jp/h25s_04/server/internal/repository"
	"github.com/traP-jp/h25s_04/server/pkg/config"
)

type Handler struct {
	repo       *repository.Repository
	traqClient *traq.APIClient
	traqAuth   context.Context
}

func New(repo *repository.Repository) *Handler {
	client, auth := config.Traq()
	return &Handler{
		repo:       repo,
		traqClient: client,
		traqAuth:   auth,
	}
}

func getUserID(xForwardedUser *string) string {
	if xForwardedUser != nil {
		return *xForwardedUser
	}

	return "traP"
}
