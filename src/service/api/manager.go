package api

import (
	"context"
	"github.com/Foundation-13/mwarehouse/src/service/utils"
	"io"

	"github.com/Foundation-13/mwarehouse/src/service/storage"
)

//go:generate mockery -name Manager -outpkg apimocks -output ./apimocks -dir .
type Manager interface {
	UploadMedia(ctx context.Context, r io.Reader, fileName string) (string, error)
}

func NewManager(stg storage.Client, idGen utils.IDGen) Manager {
	return &manager{
		stg: stg,
		idGen: idGen,
	}
}

// impl

type manager struct {
	stg storage.Client
	idGen utils.IDGen
}

func (m *manager) UploadMedia(ctx context.Context, r io.Reader, fileName string) (string, error) {
	newID := m.idGen.NewID()

	err := m.stg.Put(ctx, r, newID)
	if err != nil {
		return "", err
	}

	return newID, err
}
