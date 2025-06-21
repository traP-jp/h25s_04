package repository

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"

	"github.com/traP-jp/h25s_04/server/pkg/config"
)

type (
	Image struct {
		ID       uuid.UUID `db:"id"`
		ReviewID uuid.UUID `db:"review_id"`
	}
)

func (r *Repository) UploadImage(ctx context.Context, image io.ReadSeeker) (uuid.UUID, error) {
	imageID := uuid.New()
	// オブジェクトストレージに画像をアップロードする
	if _, err := r.client.PutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket:      aws.String(config.BucketName()),
			Key:         aws.String(imageID.String()),
			Body:        image,
			ContentType: aws.String("image/png"),
		},
	); err != nil {
		return uuid.Nil, err
	}

	return imageID, nil
}

// ファイル本体、Content-Type、エラーを返す
func (r *Repository) GetImage(ctx context.Context, imageID uuid.UUID) (io.ReadCloser, string, error) {
	result, err := r.client.GetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(config.BucketName()),
			Key:    aws.String(imageID.String()),
		},
	)
	if err != nil {
		return nil, "", err
	}

	return result.Body, *result.ContentType, nil
}
