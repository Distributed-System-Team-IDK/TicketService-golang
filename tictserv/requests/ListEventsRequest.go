package requests

import (
	"distributed.org/tictserv"
	"distributed.org/tictserv/responses"
	"github.com/gin-gonic/gin"
	"sync"
)

type ListEventsRequest struct {
	tictserv.RequestImp `json:"-"`
	ContextHolder       `json:"-"`
}

func (rq *ListEventsRequest) Exec(ts *tictserv.TicketService) tictserv.ResponseImp {
	events := ts.ListEvents()

	return &responses.ListEventsResponse{
		Status: 200,
		Events: events,
	}
}
func (rq *ListEventsRequest) GetContext() *gin.Context {
	return rq.ContextHolder.Context
}

func (rq *ListEventsRequest) GetWaitGroup() *sync.WaitGroup {
	return rq.ContextHolder.WaitGroup
}
