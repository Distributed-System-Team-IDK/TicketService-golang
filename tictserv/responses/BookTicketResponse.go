package responses

import (
	"distributed.org/tictserv"
)

type BookTicketResponse struct {
	tictserv.ResponseImp `json:"-"`
	Status               int
	Message              string
	TicketIDs            []string
}

func (rs *BookTicketResponse) GetStatus() int {
	return rs.Status
}
