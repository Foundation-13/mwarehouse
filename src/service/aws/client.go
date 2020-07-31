package aws

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Client struct {
	S3     S3Uploader
	Dynamo DynamoWrapper
}

func NewClient(region string) (*Client, error) {
	config := &aws.Config{
		Region: aws.String(region),
		MaxRetries:                    aws.Int(1),
		CredentialsChainVerboseErrors: aws.Bool(true),
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}

	session, err := session.NewSession(config)

	if err != nil {
		return nil, fmt.Errorf("failed to open AWS session: %w", err)
	}

	s3 := s3manager.NewUploader(session)
	dynamo := dynamodb.New(session)

	return &Client{
		S3:     s3,
		Dynamo: dynamo,
	}, nil
}
