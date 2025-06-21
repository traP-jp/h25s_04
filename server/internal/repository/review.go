package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	Review struct {
		Id       uuid.UUID `db:"id"`
		EateryID uuid.UUID `db:"eatery_id"`
		UserID   string    `db:"user_id"`
		Content  string    `db:"content"`
	}
)

func (r *Repository) GetEateryEateryIDReviews(ctx context.Context, eateryID uuid.UUID) ([]*Review, error) {
	reviews := []*Review{}
	if err := r.db.SelectContext(ctx, &reviews, "SELECT * FROM reviews WHERE eatery_id = ?", eateryID); err != nil {
		return nil, fmt.Errorf("select eatery: %w", err)
	}

	return reviews, nil
}
