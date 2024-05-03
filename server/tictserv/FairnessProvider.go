package tictserv

import (
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

type FairnessProvider struct {
	table       sync.Map
	maxAccepted int
}

func NewFairnessProvider(maxAccepted int) *FairnessProvider {
	return &FairnessProvider{
		table:       sync.Map{},
		maxAccepted: maxAccepted,
	}
}

func (fp *FairnessProvider) Acquire(ctx *gin.Context) bool {
	host := ctx.Request.Host
	entry, ok := fp.table.Load(host)
	var sem *Semaphore
	if !ok {
		sem = NewSemaphore(fp.maxAccepted)
		log.Printf("Fairness: new entry added")
	} else {
		sem = entry.(*Semaphore)
	}
	err := sem.Acquire()
	fp.table.Store(host, sem)
	if err != nil {
		log.Printf("Fairness: entry couldn't be acquired")
		return false
	} else {
		log.Printf("Fairness: entry acquired")
		return true
	}
}

func (fp *FairnessProvider) Release(ctx *gin.Context) bool {
	host := ctx.Request.Host
	entry, ok := fp.table.Load(host)
	var sem *Semaphore
	if !ok {
		log.Printf("Fairness: no entry to release")
	}
	sem = entry.(*Semaphore)
	err := sem.Release()
	if sem.IsEmpty() {
		log.Printf("Fairness: entry dropped")
	} else {
		fp.table.Store(host, sem)
	}
	if err != nil {
		log.Printf("Fairness: entry couldn't be released")
		return false
	} else {
		log.Printf("Fairness: entry released")
		return true
	}
}
