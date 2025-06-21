package storage

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Setup(cfg *aws.Config, optFn func(*s3.Options)) *s3.Client {
	return s3.NewFromConfig(*cfg, optFn)
}
