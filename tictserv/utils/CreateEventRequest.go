package utils

import (
	"log"
	"time"
)

type CreateEventRequest struct {
	RequestImp
	Request
	Name         string
	Date         time.Time
	TotalTickets int
}

func (rq *CreateEventRequest) Exec(ts *TicketService) (ResponseImp, error) {
	if event, err := ts.CreateEvent(rq.Name, rq.Date, rq.TotalTickets); err != nil {
		//todo: error handling
	} else {
		log.Println(event.Name, "created!") 
	}
	// todo: build response

	// response := EventResponse{
    //     EventName: event.Name,
    //     Message:   "Event created successfully!",
    // }

	return Response{}, nil
}
