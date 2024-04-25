package utils

import (
	"time"
)

type CreateEventRequest struct {
	RequestImp
	Request
	Name         string
	Date         time.Time
	TotalTickets int
}

func (r *CreateEventRequest) Handle(service *TicketService) (interface{}, error) {
	if event, err := service.CreateEvent(r.Name, r.Date, r.TotalTickets); err != nil {
		return nil, err
	} else {
		return event, nil
	}
}
