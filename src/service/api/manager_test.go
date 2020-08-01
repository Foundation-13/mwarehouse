package api_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Foundation-13/mwarehouse/src/service/api"
	"github.com/Foundation-13/mwarehouse/src/service/db/dbmocks"
	"github.com/Foundation-13/mwarehouse/src/service/storage/storagemocks"
	"github.com/Foundation-13/mwarehouse/src/service/types"
	"github.com/Foundation-13/mwarehouse/src/service/utils/utilsmocks"
)

type mocks_ struct {
	stg		*storagemocks.Client
	db		*dbmocks.Client
	idGen	*utilsmocks.IDGen
}

func newManagerWithMocks() (api.Manager, mocks_) {
	mocks := mocks_{
		stg: &storagemocks.Client{},
		db: &dbmocks.Client{},
		idGen: &utilsmocks.IDGen{},
	}

	newManager := api.NewManager(mocks.stg, mocks.db, mocks.idGen)

	return newManager, mocks
}

func TestManager_UploadMedia(t *testing.T) {
	t.Run("succeeded", func(t *testing.T) {
		subj, m := newManagerWithMocks()

		m.idGen.On("NewID").Return("123")
		m.db.On("CreateJob", mock.Anything, mock.Anything, mock.Anything).Return(&types.Job{}, nil)
		m.stg.On("Put", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		id, err := subj.UploadMedia(context.Background(), bytes.NewBuffer([]byte("")), "file.png")

		assert.NoError(t, err)
		assert.Equal(t, "123", id)

		m.stg.AssertCalled(t, "Put", mock.Anything, mock.Anything, "123")
		m.db.AssertCalled(t, "CreateJob", mock.Anything, "123", "file.png")
	})
}

func TestManager_GetJobStatus(t *testing.T) {
	t.Run("succeeded", func(t *testing.T) {
		subj, m := newManagerWithMocks()

		m.db.On("GetJobStatus", mock.Anything, mock.Anything).Return(types.JobStatus(0), nil)

		status, err := subj.GetJobStatus(context.Background(), "123")

		assert.NoError(t, err)
		assert.Equal(t, types.JobStatus(0), status)

		m.db.AssertCalled(t, "GetJobStatus", mock.Anything, "123")
	})

	t.Run("database returns an error", func(t *testing.T) {
		subj, m := newManagerWithMocks()

		m.db.On("GetJobStatus", mock.Anything, mock.Anything).
			Return(types.JobStatus(404), fmt.Errorf(""))

		status, err := subj.GetJobStatus(context.Background(), "123")

		assert.Error(t, err)
		assert.Equal(t, types.JobStatus(404), status)
	})
}