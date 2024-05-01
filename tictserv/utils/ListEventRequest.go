package utils

import (
	"log"
	"time"
)

type ListEventRequest struct {
	RequestImp
	Request
}

func (rq *ListEventRequest) Exec(ts *TicketService) (ResponseImp, error) {
	events = ts.ListEvents();
	if len(events) == 0{
		log.Println("No events available!")
	} else {
		for event := range events{
			log.Println(event.Name)
		}
	}
	// todo: build response
	return Response{}, nil
}
