package stg

import (
	"github.com/Foundation-13/mwarehouse/src/service/aws"
)

type Client interface {
	Name() string

}

func NewAWSClient(bucket string, wrapper aws.S3Wrapper) Client {
	return &s3Impl{
		wrapper:  wrapper,
		basePath: bucket,
	}
}
