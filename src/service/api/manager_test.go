package api_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Foundation-13/mwarehouse/src/service/api"
	"github.com/Foundation-13/mwarehouse/src/service/db/dbmocks"
	"github.com/Foundation-13/mwarehouse/src/service/storage/storagemocks"
	"github.com/Foundation-13/mwarehouse/src/service/types"
	"github.com/Foundation-13/mwarehouse/src/service/utils/utilsmocks"
)

func TestManager_UploadMedia(t *testing.T) {
	stg := &storagemocks.Client{}
	db := &dbmocks.Client{}
	idGen := &utilsmocks.IDGen{}

	idGen.On("NewID").Return("123")
	db.On("CreateJob", mock.Anything, mock.Anything, mock.Anything).Return(&types.Job{}, nil)
	stg.On("Put", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	subj := api.NewManager(stg, db, idGen)

	id, err := subj.UploadMedia(context.Background(), bytes.NewBuffer([]byte("")), "file.png")

	assert.NoError(t, err)
	assert.Equal(t, "123", id)

	stg.AssertCalled(t, "Put", mock.Anything, mock.Anything, "123")
	db.AssertCalled(t, "CreateJob", mock.Anything, "123", "file.png")
}
