package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	Eatery struct {
		ID          uuid.UUID `db:"id"`
		Name        string    `db:"name"`
		Description *string   `db:"description"`
	}
)

func (r *Repository) GetEatery(ctx context.Context, eateryID uuid.UUID) (*Eatery, error) {
	eatery := &Eatery{}
	if err := r.db.GetContext(ctx, eatery, "SELECT * FROM eateries WHERE id = ?", eateryID); err != nil {
		return nil, fmt.Errorf("select eatery: %w", err)
	}

	return eatery, nil
}
