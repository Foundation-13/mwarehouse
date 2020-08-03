package aws

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//go:generate mockery -name S3Uploader -outpkg awsmocks -output ./awsmocks -dir .
type S3Uploader interface {
	Upload(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

//go:generate mockery -name DynamoWrapper -outpkg awsmocks -output ./awsmocks -dir .
type DynamoWrapper interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}
