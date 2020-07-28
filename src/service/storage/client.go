package storage

import (
	"context"
	"io"

	"github.com/Foundation-13/mwarehouse/src/service/aws"
)

type Client interface {
	Name() string
	Put(ctx context.Context, r io.ReadSeeker, name string) error
}

func NewAWSClient(bucket string, wrapper aws.S3Wrapper) Client {
	return &s3Impl{
		wrapper:  wrapper,
		basePath: bucket,
	}
}
