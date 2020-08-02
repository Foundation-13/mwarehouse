package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Foundation-13/mwarehouse/src/service/config"
)

func TestFromEnvironment(t *testing.T) {
	os.Setenv("MW_AWS_REGION", "eu-central-1")
	os.Setenv("MW_TEMP_BUCKET_NAME", "test-bucket-name")
	os.Setenv("MW_LOCAL_RUN", "true")

	actualCfg, err := config.FromEnvironment()
	assert.NoError(t, err)

	expectedCfg := config.Config{
		Region:         "eu-central-1",
		TempBucketName: "test-bucket-name",
		LocalRun:		 true,
	}
	assert.Equal(t, expectedCfg, actualCfg)

	os.Clearenv()
}
