package storage

import (
	"context"
	"io"

	"github.com/Foundation-13/mwarehouse/src/service/aws"
)

//go:generate mockery -name Client -outpkg storagemocks -output ./storagemocks -dir .
type Client interface {
	Name() string

	Put(ctx context.Context, r io.Reader, key string) error
}

func NewAWSClient(bucket string, uploader aws.S3Uploader) Client {
	return &s3Impl{
		uploader: uploader,
		bucket:   bucket,
	}
}
