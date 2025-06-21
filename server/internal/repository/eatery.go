package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	// users table
	Eatery struct {
		ID          uuid.UUID `db:"id"`
		Name        string    `db:"name"`
		Description *string   `db:"description"`
	}

	CreateEateryParams struct {
		Name        string
		Description string
	}
)

func (r *Repository) CreateEateries(ctx context.Context, params CreateEateryParams) (uuid.UUID, error) {
	eateryID := uuid.New()
	if _, err := r.db.ExecContext(ctx, "INSERT INTO eateries (id, name, description) VALUES (?, ?, ?)", eateryID, params.Name, params.Description); err != nil {
		return uuid.Nil, fmt.Errorf("insert eatery: %w", err)
	}
	return eateryID, nil
}

func (r *Repository) GetEatery(ctx context.Context, eateryID uuid.UUID) (*Eatery, error) {
	eatery := &Eatery{}
	if err := r.db.GetContext(ctx, eatery, "SELECT * FROM eateries WHERE id = ?", eateryID); err != nil {
		return nil, fmt.Errorf("select eatery: %w", err)
	}

	return eatery, nil
}
