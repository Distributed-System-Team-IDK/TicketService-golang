package tictserv

import (
	"errors"
	"sync"
)

type Semaphore struct {
	count int
	max   int
	lock  sync.Mutex
}

func NewSemaphore(max int) *Semaphore {
	return &Semaphore{
		count: 0,
		max:   max,
	}
}

func (sem *Semaphore) Acquire() error {
	if sem.count < sem.max {
		sem.lock.Lock()
		sem.count += 1
		sem.lock.Unlock()
		return nil
	} else {
		return errors.New("semaphore is full")
	}
}

func (sem *Semaphore) Release() error {
	if sem.count > 0 {
		sem.lock.Lock()
		sem.count -= 1
		sem.lock.Unlock()
		return nil
	} else {
		return errors.New("semaphore is empty")
	}
}

func (sem *Semaphore) IsEmpty() bool {
	return sem.count == 0
}
