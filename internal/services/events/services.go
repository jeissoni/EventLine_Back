package events

import "github.com/jeissoni/EventLine/internal/ports"

type Service struct {
	Repository ports.EventRepository
}
