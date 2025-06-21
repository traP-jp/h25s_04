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

// そのIDのレビューが存在するかチェック
func (r *Repository) ReviewExists(ctx context.Context, reviewID uuid.UUID) (uuid.UUID, error) {
	var exists bool
	if err := r.db.GetContext(ctx, &exists, "SELECT EXISTS(SELECT 1 FROM reviews WHERE id = ?)", reviewID); err != nil {
		return uuid.Nil, fmt.Errorf("check review exists: %w", err)
	}
	if !exists {
		return uuid.Nil, fmt.Errorf("review with id %s does not exist", reviewID)
	}
	return reviewID, nil
}

func (r *Repository) DeleteReview(ctx context.Context, reviewID uuid.UUID) error {
	if _, err := r.db.ExecContext(ctx, "DELETE FROM reviews WHERE id = ?", reviewID); err != nil {
		return fmt.Errorf("delete review: %w", err)
	}
	return nil
}
