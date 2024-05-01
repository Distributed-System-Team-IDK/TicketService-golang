package requests

import (
	"distributed.org/tictsrv/src"
	"distributed.org/tictsrv/src/responses"
	"log"
)

type BookTicketRequest struct {
	src.RequestImp
	eventID    string
	numTickets int
}

func (rq *BookTicketRequest) Exec(ts *src.TicketService) src.ResponseImp {
	if ticketIDs, err := ts.BookTickets(rq.eventID, rq.numTickets); err != nil {
		return responses.ErrorResponse{
			Status:  500,
			Message: "book ticket failed: " + err.Error(),
		}
	} else {
		for ticketID := range ticketIDs {
			log.Println(ticketID, "booked")
		}
		return responses.BookTicketResponse{
			Status:    200,
			Message:   "tickets were booked successfully",
			TicketIDs: ticketIDs,
		}
	}
}
