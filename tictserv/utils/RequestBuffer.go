package utils

import (
	"log"
)

type RequestBuffer struct {
	Service *TicketService
}

func (rb *RequestBuffer) Listen(ch <-chan RequestImp) error {
	for req := range ch {
		req := req // because when for-loop continues, req changes
		go func() {
			if res, err := req.Handle(rb.Service); err != nil {
				return
			} else {
				log.Println(res)
			}
		}()
	}
	return nil
}
