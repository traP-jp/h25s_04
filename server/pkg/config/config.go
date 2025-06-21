package config

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-sql-driver/mysql"
)

func getEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return v
}

func AppAddr() string {
	return getEnv("APP_ADDR", ":8080")
}

func MySQL() *mysql.Config {
	c := mysql.NewConfig()

	c.User = getEnv("DB_USER", "root")
	c.Passwd = getEnv("DB_PASS", "pass")
	c.Net = getEnv("DB_NET", "tcp")
	c.Addr = fmt.Sprintf(
		"%s:%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
	)
	c.DBName = getEnv("DB_NAME", "app")
	c.Collation = "utf8mb4_general_ci"
	c.AllowNativePasswords = true

	return c
}

func ObjectStorage(ctx context.Context) (*aws.Config, func(*s3.Options), error) {
	accessKey := getEnv("STORAGE_ACCESS_KEY", "minioadmin")
	secret := getEnv("STORAGE_SECRET", "minioadmin")
	endpoint := getEnv("STORAGE_BASE_ENDPOINT", "http://minio:9000")
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey,
			secret,
			"",
		)),
		config.WithRegion("auto"),
	)

	if err != nil {
		return nil, nil, err
	}

	optFn := func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	}

	return &cfg, optFn, nil
}

func BucketName() string {
	return getEnv("STORAGE_BUCKET_NAME", "h25s-04")
}
