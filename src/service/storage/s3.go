package storage

import (
	"context"
	"io"

	"github.com/Foundation-13/mwarehouse/src/service/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type s3Impl struct {
	uploader aws.S3Uploader
	bucket   string
}

func (s *s3Impl) Name() string {
	return s.bucket
}

func (s *s3Impl) Put(ctx context.Context, r io.Reader, key string) error {
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Body:   r,
		Bucket: &s.bucket,
		Key:    &key,
	})

	return err
}
