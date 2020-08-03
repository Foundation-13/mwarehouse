package db

import (
	"context"
	"fmt"
	"time"

	aws2 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/Foundation-13/mwarehouse/src/service/aws"
	"github.com/Foundation-13/mwarehouse/src/service/types"
)

const (
	jobsTable = "Jobs"
)

type dynamoImpl struct {
	db aws.DynamoWrapper
}

func (d *dynamoImpl) CreateJob(ctx context.Context, key string, fileName string) (*types.Job, error) {
	job := types.Job{
		FileName: fileName,
		Key:      key,
		Created:  time.Now().UTC().UnixNano(),
		Status:   types.JobStatusCreated,
	}

	av, err := dynamodbattribute.MarshalMap(job)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal job: %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws2.String(jobsTable),
	}

	_, err = d.db.PutItem(input)
	if err != nil {
		return nil, fmt.Errorf("failed to store job: %w", err)
	}

	return &job, nil
}

func (d *dynamoImpl) GetJobStatus(ctx context.Context, key string) (types.Job, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws2.String(jobsTable),
		Key: map[string]*dynamodb.AttributeValue{
			"Key": {
				S: aws2.String(key),
			},
		},
	}

	res, err := d.db.GetItem(input)
	if err != nil {
		return types.Job{}, fmt.Errorf("failed to get item: %w", err)
	}

	job := types.Job{}
	err = dynamodbattribute.UnmarshalMap(res.Item, &job)
	if err != nil {
		return types.Job{}, fmt.Errorf("failed to unmarshal job: %w", err)
	}

	if job.Key == "" {
		return types.Job{}, fmt.Errorf("job with a key: %v not found", key)
	}

	return job, nil
}
