package requests

import (
	"distributed.org/tictserv"
	"distributed.org/tictserv/responses"
	"log"
)

type ListEventRequest struct {
	tictserv.RequestImp
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

	return responses.ListEventsResponse{
		Status: 200,
		Events: events,
	}
}
