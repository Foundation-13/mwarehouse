package api

import (
	"context"
	"io"

	"github.com/Foundation-13/mwarehouse/src/service/db"
	"github.com/Foundation-13/mwarehouse/src/service/storage"
	"github.com/Foundation-13/mwarehouse/src/service/types"
	"github.com/Foundation-13/mwarehouse/src/service/utils"
)

//go:generate mockery -name Manager -outpkg apimocks -output ./apimocks -dir .
type Manager interface {
	UploadMedia(ctx context.Context, r io.Reader, fileName string) (string, error)
	GetJobStatus(ctx context.Context, key string) (types.Job, error)
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
		return "", err
	}

	err = m.stg.Put(ctx, r, newID)
	if err != nil {
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
