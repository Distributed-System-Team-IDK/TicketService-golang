package responses

import "distributed.org/tictsrv/src"

type CreateEventResponse struct {
	src.ResponseImp
	Status  int
	Message string
}
