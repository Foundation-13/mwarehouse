package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	aws_s3 "github.com/aws/aws-sdk-go/service/s3"
)

type Client struct {
	S3 S3Wrapper
}

func NewClient() (*Client, error) {
	config := &aws.Config{
		Region: aws.String("us-east-2"),
	}

	session, err := session.NewSession(config)

	if err != nil {
		return nil, fmt.Errorf("failed to open AWS session: %w", err)
	}

	s3 := aws_s3.New(session)

	return &Client{
		S3: s3,
	}, nil
}
