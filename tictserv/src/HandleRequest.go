package src

func HandleRequest(ts *TicketService, rqch <-chan RequestImp, rsch chan<- ResponseImp) {
	for req := range rqch {
		req := req
		go func() {
			res := req.Exec(ts)
			rsch <- res
		}()
	}
}
