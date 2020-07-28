package api

import (
	"context"
	"io"

	"github.com/Foundation-13/mwarehouse/src/service/storage"
)

//go:generate mockery -name Manager -outpkg apimocks -output ./apimocks -dir .
type Manager interface {
	UploadMedia(ctx context.Context, r io.ReadSeeker, fileName string, contentType string) (string, error)
}

func NewManager(stg storage.Client) Manager {
	return &manager{
		stg: stg,
	}
}

// impl

type manager struct {
	stg storage.Client
}

func (m *manager) UploadMedia(ctx context.Context, r io.ReadSeeker, fileName string, contentType string) (string, error) {
	err := m.stg.Put(ctx, r, fileName)

	return "1", err
}
