package tictserv

func HandleRequest(ts *TicketService, rqch <-chan RequestImp) {
	for req := range rqch {
		req := req
		go func() {
			res := req.Exec(ts)
			ctx := req.GetContext()
			wg := req.GetWaitGroup()

			ctx.JSON(res.GetStatus(), res)
			wg.Done()
		}()
	}
}
