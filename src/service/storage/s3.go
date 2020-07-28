package storage

import (
	"context"
	"github.com/Foundation-13/mwarehouse/src/service/aws"
	aws_s3 "github.com/aws/aws-sdk-go/service/s3"
	"io"
)

type s3Impl struct {
	wrapper  aws.S3Wrapper
	basePath string
}

func (s *s3Impl) Name() string {
	return s.basePath
}

func (s *s3Impl) Put(ctx context.Context, r io.ReadSeeker, name string) error {
	_, err := s.wrapper.PutObject(&aws_s3.PutObjectInput{
		Key: &name,
		Bucket: &s.basePath,
		Body: r,
	})

	return err
}