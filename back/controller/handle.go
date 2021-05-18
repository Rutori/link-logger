package controller

import "context"

type handle struct {
	Path   string
	Method string
	Func   func(ctx context.Context, request []byte) (response []byte, err error)
}
