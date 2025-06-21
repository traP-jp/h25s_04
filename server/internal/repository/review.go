package repository

import "github.com/google/uuid"

type (
	Review struct {
		Id       uuid.UUID `db:"id"`
		EateryID uuid.UUID `db:"eatery_id"`
		UserID   string    `db:"user_id"`
		Content  string    `db:"content"`
	}
)
