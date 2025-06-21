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

func (r *Repository) UpdateEatery(ctx context.Context, eateryID uuid.UUID, eatery Eatery, userID string) error {
	// Assuming you have a table `eatery_updates` to log updates
	query := `
		UPDATE eateries 
		SET name = ?, description = ? 
		WHERE id = ? AND user_id = ?
	`
	if _, err := r.db.ExecContext(ctx, query, eatery.Name, eatery.Description, eateryID, userID); err != nil {
		return fmt.Errorf("update eatery: %w", err)
	}

	return nil
}
