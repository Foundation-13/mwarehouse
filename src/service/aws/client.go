package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Client struct {
	S3 S3Uploader
}

func NewClient() (*Client, error) {
	config := &aws.Config{
		Region: aws.String("us-east-2"),
	}

	session, err := session.NewSession(config)

	if err != nil {
		return nil, fmt.Errorf("failed to open AWS session: %w", err)
	}

	s3 := s3manager.NewUploader(session)

	return &Client{
		S3: s3,
	}, nil
}
