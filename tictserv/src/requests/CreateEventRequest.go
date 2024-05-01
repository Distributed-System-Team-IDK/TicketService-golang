package requests

import (
	"distributed.org/tictsrv/src"
	"distributed.org/tictsrv/src/responses"
	"log"
	"time"
)

type CreateEventRequest struct {
	src.RequestImp
	Name         string
	Date         time.Time
	TotalTickets int
}

func (rq *CreateEventRequest) Exec(ts *src.TicketService) src.ResponseImp {
	if event, err := ts.CreateEvent(rq.Name, rq.Date, rq.TotalTickets); err != nil {
		return responses.ErrorResponse{
			Status:  500,
			Message: "create event failed: " + err.Error(),
		}
	} else {
		log.Println(event.Name, "created!")
		return responses.CreateEventResponse{
			Status:  200,
			Message: "event created successfully",
		}
	}
}
