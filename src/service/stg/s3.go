package stg

import (
	"github.com/Foundation-13/mwarehouse/src/service/aws"
)

type s3Impl struct {
	wrapper  aws.S3Wrapper
	basePath string
}

func (s *s3Impl) Name() string {
	return s.basePath
}