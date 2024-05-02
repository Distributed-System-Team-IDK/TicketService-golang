package responses

import (
	"distributed.org/tictserv"
)

type ListEventsResponse struct {
	tictserv.ResponseImp `json:"-"`
	Status               int
	Events               []*tictserv.Event
}

func (rs *ListEventsResponse) GetStatus() int {
	return rs.Status
}
