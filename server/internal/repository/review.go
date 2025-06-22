package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type (
	Review struct {
		Id        uuid.UUID `db:"id"`
		EateryID  uuid.UUID `db:"eatery_id"`
		UserID    string    `db:"user_id"`
		Content   string    `db:"content"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	CreateEateryReviewParams struct {
		ID       uuid.UUID `json:"id"`
		EateryID uuid.UUID `json:"eatery_id"`
		Content  string    `json:"content"`
		UserID   string    `json:"user_id"`
	}
)

func (r *Repository) PostEateryReview(ctx context.Context, params CreateEateryReviewParams) (uuid.UUID, error) {
	// パスパラメータから受け取ったeateryIDをparams.EateryIDにセットする場合は、
	// 呼び出し元のハンドラでパスパラメータを取得し、CreateEateryReviewParamsにセットしてください。
	// ここ（リポジトリ層）ではパスパラメータの取得は行いません。

	var exists bool
	if err := r.db.GetContext(ctx, &exists, "SELECT EXISTS(SELECT 1 FROM eateries WHERE id = ?)", params.EateryID); err != nil {
		return uuid.Nil, fmt.Errorf("failed to check eatery existence: %w", err)
	}
	if !exists {
		return uuid.Nil, fmt.Errorf("eatery with that id does not exist")
	}

	reviewID := uuid.New()

	if _, err := r.db.ExecContext(ctx, "INSERT INTO reviews (id, eatery_id, user_id, content) VALUES (?, ?, ?, ?)", reviewID, params.EateryID, params.UserID, params.Content); err != nil {
		return uuid.Nil, fmt.Errorf("insert eatery review: %w", err)
	}
	return reviewID, nil
}

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

func (r *Repository) GetReview(ctx context.Context, reviewID uuid.UUID) (*Review, error) {
	var review Review
	err := r.db.GetContext(ctx, &review, `SELECT * FROM reviews WHERE id = ?`, reviewID)
	if err != nil {
		return nil, fmt.Errorf("get review: %w", err)
	}
	return &review, nil
}

func (r *Repository) UpdateReviewContent(ctx context.Context, reviewID uuid.UUID, content string) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE reviews SET content = ? WHERE id = ?`, content, reviewID)
	if err != nil {
		return fmt.Errorf("update review content: %w", err)
	}
	return nil
}
