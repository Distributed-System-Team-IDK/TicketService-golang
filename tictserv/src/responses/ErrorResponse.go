package responses

import "distributed.org/tictsrv/src"

type ErrorResponse struct {
	src.ResponseImp
	Status  int
	Message string
}
