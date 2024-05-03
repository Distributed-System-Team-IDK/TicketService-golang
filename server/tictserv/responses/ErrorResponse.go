package responses

import (
	"distributed.org/tictserv"
)

type ErrorResponse struct {
	tictserv.ResponseImp `json:"-"`
	Status               int
	Message              string
}

func (rs *ErrorResponse) GetStatus() int {
	return rs.Status
}
