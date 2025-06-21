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

func (r *Repository) GetEateryEateryIDReviews(ctx context.Context, eateryID uuid.UUID, limit, offset int) ([]*Review, error) {
	reviews := []*Review{}
	const query = `
		SELECT *
		FROM reviews
		WHERE eatery_id = ?
		ORDER BY id
		LIMIT ? OFFSET ?
	`

	if err := r.db.SelectContext(ctx, &reviews, query, eateryID, limit, offset); err != nil {
		return nil, fmt.Errorf("select eatery reviews: %w", err)
	}

	return reviews, nil
}
