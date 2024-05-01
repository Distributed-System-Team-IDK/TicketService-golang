package responses

import (
	"distributed.org/tictserv"
)

type BookTicketResponse struct {
	tictserv.ResponseImp
	Status    int
	Message   string
	TicketIDs []string
}
