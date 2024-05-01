package responses

import (
	"distributed.org/tictsrv/src"
)

type ListEventsResponse struct {
	src.ResponseImp
	Status int
	Events []*src.Event
}
