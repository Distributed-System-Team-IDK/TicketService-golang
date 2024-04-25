package CA2

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

var generateUUID = func() string {
	return uuid.New().String()
}

type TicketService struct {
	events  sync.Map
	tickets sync.Map
}

func (ts *TicketService) CreateEvent(name string, date time.Time, totalTickets int) (*Event, error) {
	event := &Event{
		ID:               generateUUID(),
		Name:             name,
		Date:             date,
		TotalTickets:     totalTickets,
		AvailableTickets: totalTickets,
	}

	ts.events.Store(event.ID, event)
	return event, nil
}

func (ts *TicketService) ListEvents() []*Event {
	var events []*Event
	ts.events.Range(func(key, value interface{}) bool {
		event := value.(*Event)
		events = append(events, event)
		return true
	})
	return events
}

func (ts *TicketService) BookTickets(eventID string, numTickets int) ([]string, error) {
	event, ok := ts.events.Load(eventID)
	if !ok {
		return nil, fmt.Errorf("event not found")
	}

	var ev *Event
	if ev, ok = event.(*Event); !ok {
		return nil, fmt.Errorf("invalid event")
	}

	if ev.AvailableTickets < numTickets {
		return nil, fmt.Errorf("not enough tickets available")
	}

	var ticketIDs []string
	for i := 0; i < numTickets; i++ {
		ticket := &Ticket{
			ID:      generateUUID(),
			EventID: ev.ID,
		}
		ticketIDs = append(ticketIDs, ticket.ID)
		ts.tickets.Store(ticket.ID, ticket)
	}

	ev.mu.Lock()
	ev.AvailableTickets -= numTickets
	ev.mu.Unlock()

	ts.events.Store(eventID, ev)

	return ticketIDs, nil
}
