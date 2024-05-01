package utils

type RequestImp interface {
	Exec(ts *TicketService) (ResponseImp, error)
}
