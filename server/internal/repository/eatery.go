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
		Description *string   `db:"description"`
	}

	CreateEateryParams struct {
		Name        string
		Description *string
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

// Eateryの存在チェック
func (r *Repository) EateryExists(ctx context.Context, params CreateEateryReviewParams) (uuid.UUID, error) {
	var exists bool
	if err := r.db.GetContext(ctx, &exists, "SELECT EXISTS(SELECT 1 FROM eateries WHERE id = ?)", params.EateryID); err != nil {
		return uuid.Nil, fmt.Errorf("failed to check eatery existence: %w", err)
	}
	if !exists {
		return uuid.Nil, fmt.Errorf("eatery with that id does not exist")
	}
	return uuid.Nil, nil
}

func (r *Repository) UpdateEatery(ctx context.Context, eateryID uuid.UUID, eatery Eatery) error {
	// Assuming you have a table `eatery_updates` to log updates
	query := `
		UPDATE eateries 
		SET name = ?, description = ? 
		WHERE id = ?
	`
	if _, err := r.db.ExecContext(ctx, query, eatery.Name, eatery.Description, eateryID); err != nil {
		return fmt.Errorf("update eatery: %w", err)
	}

	return nil
}
