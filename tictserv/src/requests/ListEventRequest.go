package requests

import (
	"distributed.org/tictsrv/src"
	"distributed.org/tictsrv/src/responses"
	"log"
)

type ListEventRequest struct {
	src.RequestImp
}

func (rq *ListEventRequest) Exec(ts *src.TicketService) src.ResponseImp {
	events := ts.ListEvents()

	// log
	if len(events) == 0 {
		log.Println("No events available!")
	} else {
		for _, event := range events {
			log.Println(event.Name)
		}
	}

	return responses.ListEventsResponse{
		Status: 200,
		Events: events,
	}
}
