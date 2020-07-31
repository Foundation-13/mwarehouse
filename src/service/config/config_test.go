package config_test

import (
	"github.com/Foundation-13/mwarehouse/src/service/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFromEnvironment(t *testing.T) {
	os.Setenv("REGION", "eu-central-1")
	os.Setenv("TEMP_BUCKET_NAME", "test-bucket-name")

	expectedCfg := config.Config{
		Region:         "eu-central-1",
		TempBucketName: "test-bucket-name",
	}

	actualCfg, err := config.FromEnvironment()

	assert.NoError(t, err)
	assert.Equal(t, expectedCfg, actualCfg)

	os.Clearenv()
}