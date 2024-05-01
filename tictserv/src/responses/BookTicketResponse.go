package responses

import "distributed.org/tictsrv/src"

type BookTicketResponse struct {
	src.ResponseImp
	Status    int
	Message   string
	TicketIDs []string
}
