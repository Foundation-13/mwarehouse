package api

import (
	"context"
	"fmt"
	"io"
)

//go:generate mockery -name Manager -outpkg apimocks -output ./apimocks -dir .
type Manager interface {
	UploadMedia(ctx context.Context, r io.Reader, fileName string, contentType string) (string, error)
}

func NewManager() Manager {
	return &manager{}
}

// impl

type manager struct {
}

func (m *manager) UploadMedia(ctx context.Context, r io.Reader, fileName string, contentType string) (string, error) {
	return "", fmt.Errorf("not implemented")
}
