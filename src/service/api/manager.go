package api

import (
	"context"
	"fmt"
	"io"

	"github.com/Foundation-13/mwarehouse/src/service/db"
	"github.com/Foundation-13/mwarehouse/src/service/log"
	"github.com/Foundation-13/mwarehouse/src/service/storage"
	"github.com/Foundation-13/mwarehouse/src/service/types"
	"github.com/Foundation-13/mwarehouse/src/service/utils"
)

//go:generate mockery -name Manager -outpkg apimocks -output ./apimocks -dir .
type Manager interface {
	UploadMedia(ctx context.Context, r io.Reader, fileName string) (string, error)
	GetJobStatus(ctx context.Context, key string) (types.Job, error)
	ProcessMedia(ctx context.Context, key string, filters types.Filters) error
}

func NewManager(stg storage.Client, db db.Client, idGen utils.IDGen) Manager {
	return &manager{
		stg:   stg,
		db:    db,
		idGen: idGen,
	}
}

// impl

type manager struct {
	stg   storage.Client
	db    db.Client
	idGen utils.IDGen
}

func (m *manager) UploadMedia(ctx context.Context, r io.Reader, fileName string) (string, error) {
	newID := m.idGen.NewID()

	_, err := m.db.CreateJob(ctx, newID, fileName)
	if err != nil {
		err = fmt.Errorf("failed to create job, %w", err)
		log.FromContext(ctx).WithError(err).Error("failed to upload media")
		return "", err
	}

	err = m.stg.Put(ctx, r, newID)
	if err != nil {
		err = fmt.Errorf("failed to create add media into storage, %w", err)
		log.FromContext(ctx).WithError(err).Error("failed to upload media")
		return "", err
	}

	return newID, err
}

func (m *manager) GetJobStatus(ctx context.Context, key string) (types.Job, error) {
	res, err := m.db.GetJobStatus(ctx, key)
	if err != nil {
		return types.Job{}, err
	}

	return res, nil
}

func (m *manager) ProcessMedia(ctx context.Context, key string, filters types.Filters) error {
	return nil
}
