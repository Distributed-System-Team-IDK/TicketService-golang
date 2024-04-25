package utils

type RequestImp interface {
	Handle(service *TicketService) (interface{}, error)
	Respond() error
}
