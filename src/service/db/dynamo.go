package db

import (
	"context"
	"time"

	aws2 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/Foundation-13/mwarehouse/src/service/aws"
	"github.com/Foundation-13/mwarehouse/src/service/log"
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
		log.FromContext(ctx).WithError(err).Error("failed to marshal job")
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws2.String(jobsTable),
	}

	_, err = d.db.PutItem(input)
	if err != nil {
		log.FromContext(ctx).WithError(err).Error("failed to store job")
		return nil, err
	}

	return &job, nil
}
