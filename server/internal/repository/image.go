package repository

import (
	"context"
	"io"

	"github.com/google/uuid"
)

type (
	Image struct {
		ID       uuid.UUID `db:"id"`
		ReviewID uuid.UUID `db:"review_id"`
	}
)

func (r *Repository) UploadImage(ctx context.Context, image io.Reader) (uuid.UUID, error) {
	imageID := uuid.New()
	// オブジェクトストレージに画像をアップロードする
	return imageID, nil
}
