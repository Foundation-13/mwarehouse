package api

import (
	"context"
	"fmt"
	"io"

	"github.com/Foundation-13/mwarehouse/src/service/stg"
)

//go:generate mockery -name Manager -outpkg apimocks -output ./apimocks -dir .
type Manager interface {
	UploadMedia(ctx context.Context, r io.Reader, fileName string, contentType string) (string, error)
}

func NewManager(stg stg.Client) Manager {
	return &manager{
		stg: stg,
	}
}

// impl

type manager struct {
	stg stg.Client
}

func (m *manager) UploadMedia(ctx context.Context, r io.Reader, fileName string, contentType string) (string, error) {
	return "", fmt.Errorf("not implemented")
}
