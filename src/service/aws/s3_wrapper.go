package aws

import aws_s3 "github.com/aws/aws-sdk-go/service/s3"

//go:generate mockery -name Wrapper -outpkg awsmocks -output ./awsmocks -dir .
type S3Wrapper interface {
	GetObject(*aws_s3.GetObjectInput) (*aws_s3.GetObjectOutput, error)
	PutObject(*aws_s3.PutObjectInput) (*aws_s3.PutObjectOutput, error)
	DeleteObject(*aws_s3.DeleteObjectInput) (*aws_s3.DeleteObjectOutput, error)
}
