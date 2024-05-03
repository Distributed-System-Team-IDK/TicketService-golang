package responses

import (
	"distributed.org/tictserv"
)

type CreateEventResponse struct {
	tictserv.ResponseImp `json:"-"`
	Status               int
	Message              string
	ID                   string
}

func (rs *CreateEventResponse) GetStatus() int {
	return rs.Status
}
