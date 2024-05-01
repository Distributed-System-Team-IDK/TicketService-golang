package responses

import (
	"distributed.org/tictserv"
)

type CreateEventResponse struct {
	tictserv.ResponseImp
	Status  int
	Message string
}
