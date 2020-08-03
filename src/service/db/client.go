package db

import (
	"context"
	"github.com/Foundation-13/mwarehouse/src/service/aws"
	"github.com/Foundation-13/mwarehouse/src/service/types"
)

//go:generate mockery -name Client -outpkg dbmocks -output ./dbmocks -dir .
type Client interface {
	CreateJob(ctx context.Context, key string, fileName string) (*types.Job, error)
	GetJobStatus(ctx context.Context, key string) (types.Job, error)
}

func NewDynamoDBClient(wrapper aws.DynamoWrapper) Client {
	return &dynamoImpl{db: wrapper}
}