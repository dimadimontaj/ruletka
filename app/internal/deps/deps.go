package deps

import (
	"context"
)

//go:generate mockgen -source deps.go -destination=mocks/deps.go -package mocks

type Logger interface {
	Info(ctx context.Context, message string, args ...any)
	Error(ctx context.Context, err error, args ...any)
}