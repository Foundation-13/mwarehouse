package config

import (
	"fmt"
	"os"
)

type Config struct {
	Region 			string
	TempBucketName 	string
}

func FromEnvironment () (Config, error) {
	region := os.Getenv("REGION")
	if region == "" {
		return Config{}, fmt.Errorf("REGION env variable is not defined")
	}

	tempBucketName := os.Getenv("TEMP_BUCKET_NAME")
	if tempBucketName == "" {
		return Config{}, fmt.Errorf("TEMP_BUCKEY_NAME env variable is not defined")
	}

	return Config{
		Region: region,
		TempBucketName: tempBucketName,
	}, nil
}
