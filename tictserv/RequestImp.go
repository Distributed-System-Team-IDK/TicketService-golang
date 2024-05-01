package tictserv

type RequestImp interface {
	Exec(ts *TicketService) ResponseImp
}
