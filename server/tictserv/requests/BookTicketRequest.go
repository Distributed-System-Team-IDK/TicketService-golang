package requests

import (
	"distributed.org/tictserv"
	"distributed.org/tictserv/responses"
	"github.com/gin-gonic/gin"
	"sync"
)

type BookTicketRequest struct {
	tictserv.RequestImp `json:"-"`
	ContextHolder       `json:"-"`
	EventID             string `json:"event-id"`
	NumTickets          int    `json:"num-tickets"`
}

func (rq *BookTicketRequest) Exec(ts *tictserv.TicketService) tictserv.ResponseImp {
	if ticketIDs, err := ts.BookTickets(rq.EventID, rq.NumTickets); err != nil {
		return &responses.ErrorResponse{
			Status:  500,
			Message: "book ticket failed: " + err.Error(),
		}
	} else {
		return &responses.BookTicketResponse{
			Status:    200,
			Message:   "tickets were booked successfully",
			TicketIDs: ticketIDs,
		}
	}
}

func (rq *BookTicketRequest) GetContext() *gin.Context {
	return rq.ContextHolder.Context
}

func (rq *BookTicketRequest) GetWaitGroup() *sync.WaitGroup {
	return rq.ContextHolder.WaitGroup
}
