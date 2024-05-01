package utils

import (
	"log"
	"time"
)

type BookTicketRequest struct {
	RequestImp
	Request
	eventID string
	numTickets int
}

func (rq *BookTicketRequest) Exec(ts *TicketService) (ResponseImp, error) {
	if ticketIDs, err := ts.BookTickets(rq.eventID, rq.numTickets); err != nil {
		//todo: error handling
	} else {
		for ticketID := range ticketIDs{
			log.Println(ticketID, "booked")
		}
	}
	// todo: build response
	return Response{}, nil
}
