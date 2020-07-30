package config

import (
	"fmt"
	"os"
)

type Config struct {
	TestMode		string
	BucketName 		string
	Region 			string
}

func FromEnvironment() (Config, error) {
	testMode := os.Getenv("TEST_MODE")

	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		return Config{}, fmt.Errorf("BUCKET_NAME env variable is not defined")
	}

	region := os.Getenv("REGION")
	if region == "" {
		return Config{}, fmt.Errorf("REGION env variable is not defined")
	}

	return Config{
		TestMode: testMode,
		BucketName: bucketName,
		Region: region,
	}, nil
}