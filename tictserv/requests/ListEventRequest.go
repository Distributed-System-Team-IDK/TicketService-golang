package requests

import (
	"distributed.org/tictserv"
	"distributed.org/tictserv/responses"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

type ListEventRequest struct {
	tictserv.RequestImp `json:"-"`
	ContextHolder       `json:"-"`
}

func (rq *ListEventRequest) Exec(ts *tictserv.TicketService) tictserv.ResponseImp {
	events := ts.ListEvents()

	// log
	if len(events) == 0 {
		log.Println("No events available!")
	} else {
		for _, event := range events {
			log.Println(event.Name)
		}
	}

	return &responses.ListEventsResponse{
		Status: 200,
		Events: events,
	}
}
func (rq *ListEventRequest) GetContext() *gin.Context {
	return rq.ContextHolder.Context
}

func (rq *ListEventRequest) GetWaitGroup() *sync.WaitGroup {
	return rq.ContextHolder.WaitGroup
}
