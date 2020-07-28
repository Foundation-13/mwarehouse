package api_test

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Foundation-13/mwarehouse/src/service/api"
	"github.com/Foundation-13/mwarehouse/src/service/storage/storagemocks"
	"github.com/Foundation-13/mwarehouse/src/service/utils/utilsmocks"
)

func TestManager_UploadMedia(t *testing.T) {
	stg := &storagemocks.Client{}
	idGen := &utilsmocks.IDGen{}

	stg.On("Put", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	idGen.On("NewID").Return("123")

	subj := api.NewManager(stg, idGen)

	id, err := subj.UploadMedia(context.Background(), bytes.NewBuffer([]byte("")), "file.png")

	assert.NoError(t, err)
	assert.Equal(t, "123", id)

	stg.AssertCalled(t, "Put", mock.Anything, mock.Anything, "123")
}



