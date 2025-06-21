package repository

import (
	"io"

	"github.com/google/uuid"
)

type (
	Image struct {
		ID       uuid.UUID `db:"id"`
		ReviewID uuid.UUID `db:"review_id"`
	}
)

func (r *Repository) UploadImage(reviewID uuid.UUID, image io.Reader) (uuid.UUID, error) {
	imageID := uuid.New()
	// オブジェクトストレージへにアップロードし、データベースにimageIDとreviewIDを保存する
	return imageID, nil
}
