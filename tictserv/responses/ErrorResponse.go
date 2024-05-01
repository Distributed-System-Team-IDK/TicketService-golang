package responses

import (
	"distributed.org/tictserv"
)

type ErrorResponse struct {
	tictserv.ResponseImp
	Status  int
	Message string
}
