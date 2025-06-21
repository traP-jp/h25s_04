package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	Eatery struct {
		ID          uuid.UUID `db:"id"` // UUID
		Name        string    `db:"name"`
		Description string    `db:"description"`
	}
)

func (r *Repository) GetEateries(ctx context.Context) ([]*Eatery, error) {
	eateries := []*Eatery{}
	if err := r.db.SelectContext(ctx, &eateries, "SELECT * FROM eateries"); err != nil {
		return nil, fmt.Errorf("failed to get eateries: %w", err)
	}
	return eateries, nil
}
