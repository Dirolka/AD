package Tickets

import "fmt"

type Ticket struct {
	ID          string
	Title       string
	Description string
	Priority    int // 1 (low), 2 (medium), 3 (high)
	AssigneeID  string
	Status      string // "OPEN" or "DONE"
}

type TicketStore struct {
	items map[string]Ticket
}

func NewTicketStore() *TicketStore {
	return &TicketStore{items: make(map[string]Ticket)}
}

func (s *TicketStore) Create(t Ticket) error {
	if t.ID == "" {
		return fmt.Errorf("ticket id must be non-empty")
	}
	if _, exists := s.items[t.ID]; exists {
		return fmt.Errorf("ticket id must be unique")
	}
	if t.Title == "" {
		return fmt.Errorf("ticket title must be non-empty")
	}
	if t.Priority < 1 || t.Priority > 3 {
		return fmt.Errorf("ticket priority must be 1..3")
	}
	if t.Status != "OPEN" {
		return fmt.Errorf("ticket status must start as OPEN")
	}

	s.items[t.ID] = t
	return nil
}

func (s *TicketStore) Assign(ticketID string, assigneeID string) error {
	if assigneeID == "" {
		return fmt.Errorf("assignee id must be non-empty")
	}

	t, ok := s.items[ticketID]
	if !ok {
		return fmt.Errorf("ticket does not exist")
	}
	if t.Status != "OPEN" {
		return fmt.Errorf("ticket must be OPEN")
	}

	t.AssigneeID = assigneeID
	s.items[ticketID] = t
	return nil
}

func (s *TicketStore) Resolve(ticketID string) error {
	t, ok := s.items[ticketID]
	if !ok {
		return fmt.Errorf("ticket does not exist")
	}

	t.Status = "DONE"
	s.items[ticketID] = t
	return nil
}

func (s *TicketStore) ListAll() []Ticket {
	out := make([]Ticket, 0, len(s.items))
	for _, t := range s.items {
		out = append(out, t)
	}
	return out
}

func (s *TicketStore) ListByStatus(status string) []Ticket {
	out := make([]Ticket, 0)
	for _, t := range s.items {
		if t.Status == status {
			out = append(out, t)
		}
	}
	return out
}

func (s *TicketStore) ListUnassigned() []Ticket {
	out := make([]Ticket, 0)
	for _, t := range s.items {
		if t.Status == "OPEN" && t.AssigneeID == "" {
			out = append(out, t)
		}
	}
	return out
}
