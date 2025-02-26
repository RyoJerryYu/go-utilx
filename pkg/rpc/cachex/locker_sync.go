package cachex

import (
	"context"
	"sync"
)

type syncLockObtainer struct {
	lock *sync.Mutex
}

func NewSyncLockObtainer() LockObtainer {
	return &syncLockObtainer{
		lock: &sync.Mutex{},
	}
}

func (l *syncLockObtainer) Obtain(ctx context.Context) (Lock, error) {
	l.lock.Lock()
	return &syncLock{lock: l.lock}, nil
}

type syncLock struct {
	lock *sync.Mutex
}

func (l *syncLock) Release(ctx context.Context) error {
	l.lock.Unlock()
	return nil
}
