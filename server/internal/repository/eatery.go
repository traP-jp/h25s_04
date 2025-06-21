package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_04/server/internal/schema"
)

type (
	Eatery struct {
		ID          uuid.UUID `db:"id"` // UUID
		Name        string    `db:"name"`
		Description string    `db:"description"`
	}
)

func (r *Repository) GetEateries(ctx context.Context, params schema.GetEateriesParams) ([]*Eatery, error) {
	eateries := []*Eatery{}
	if params.Query != nil {
		query := "%" + *params.Query + "%"
		if err := r.db.SelectContext(ctx, &eateries, "SELECT * FROM eateries WHERE name LIKE ? ", query); err != nil {
			return nil, fmt.Errorf("failed to filter eateries by query %q: %w", *params.Query, err)
		}
	}
	// If no query is provided, fetch all eateries
	if params.Query == nil {
		if err := r.db.SelectContext(ctx, &eateries, "SELECT * FROM eateries"); err != nil {
			return nil, fmt.Errorf("failed to fetch all eateries: %w", err)
		}
	}
	return eateries, nil
}
