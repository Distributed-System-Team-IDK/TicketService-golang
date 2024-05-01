package src

type RequestImp interface {
	Exec(ts *TicketService) ResponseImp
}
