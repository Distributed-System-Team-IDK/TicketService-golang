package requests

import (
	"distributed.org/tictserv"
	"distributed.org/tictserv/responses"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type CreateEventRequest struct {
	tictserv.RequestImp `json:"-"`
	ContextHolder       `json:"-"`
	Name                string
	Date                time.Time
	TotalTickets        int
}

func (rq *CreateEventRequest) Exec(ts *tictserv.TicketService) tictserv.ResponseImp {
	if event, err := ts.CreateEvent(rq.Name, rq.Date, rq.TotalTickets); err != nil {
		return &responses.ErrorResponse{
			Status:  500,
			Message: "create event failed: " + err.Error(),
		}
	} else {
		return &responses.CreateEventResponse{
			Status:  200,
			ID:      event.ID,
			Message: "event created successfully",
		}
	}
}

func (rq *CreateEventRequest) GetContext() *gin.Context {
	return rq.ContextHolder.Context
}

func (rq *CreateEventRequest) GetWaitGroup() *sync.WaitGroup {
	return rq.ContextHolder.WaitGroup
}
