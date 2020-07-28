package utils

import "github.com/rs/xid"

//go:generate mockery -name IDGen -outpkg utilsmocks -output ./utilsmocks -dir .
type IDGen interface {
	NewID() string
}

type XID struct {}

func (XID) NewID() string {
	return xid.New().String()
}

