package utils

func HandleRequest(ts *TicketService, rqch <-chan RequestImp, rsch chan<- ResponseImp) {
	for req := range rqch {
		req := req
		go func() {
			if res, err := req.Exec(ts); err != nil {
				//todo: error handling
			} else {
				rsch <- res
			}
		}()
	}
}
