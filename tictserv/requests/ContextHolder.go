package requests

import (
	"distributed.org/tictserv"
	"github.com/gin-gonic/gin"
	"sync"
)

type ContextHolder struct {
	tictserv.RequestImp

	Context   *gin.Context
	WaitGroup *sync.WaitGroup
}
