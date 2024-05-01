package responses

import (
	"distributed.org/tictserv"
)

type ListEventsResponse struct {
	tictserv.ResponseImp
	Status int
	Events []*tictserv.Event
}
