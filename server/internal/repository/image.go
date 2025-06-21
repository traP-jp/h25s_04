package repository

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	if _, err := r.client.PutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket:      aws.String("leaq"),
			Key:         aws.String(imageID.String()),
			Body:        image,
			ContentType: aws.String("image/png"),
		},
	); err != nil {
		return uuid.Nil, err
	}

	return imageID, nil
}
