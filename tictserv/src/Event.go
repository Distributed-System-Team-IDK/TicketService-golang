package src

import (
	"sync"
	"time"
)

type Event struct {
	mu               sync.Mutex
	ID               string
	Name             string
	Date             time.Time
	TotalTickets     int
	AvailableTickets int
}
