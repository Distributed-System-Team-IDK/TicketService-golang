package tictserv

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type RequestImp interface {
	Exec(ts *TicketService) ResponseImp
	GetContext() *gin.Context
	GetWaitGroup() *sync.WaitGroup
}
