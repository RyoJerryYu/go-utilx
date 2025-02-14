package loggerx

import "context"

type Loggerf interface {
	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
}

type NoopLoggerf struct{}

func (l NoopLoggerf) Debugf(ctx context.Context, format string, args ...interface{}) {}
func (l NoopLoggerf) Infof(ctx context.Context, format string, args ...interface{})  {}
func (l NoopLoggerf) Warnf(ctx context.Context, format string, args ...interface{})  {}
func (l NoopLoggerf) Errorf(ctx context.Context, format string, args ...interface{}) {}
