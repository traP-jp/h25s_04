package handler

import (
	"github.com/traP-jp/h25s_04/server/internal/repository"
)

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func getUserID(xForwardedUser *string) string {
	if xForwardedUser != nil {
		return *xForwardedUser
	}

	return "traP"
}
