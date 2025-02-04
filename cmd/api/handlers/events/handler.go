package events

import "github.com/jeissoni/EventLine/internal/ports"

type EventHandler struct {
	EventService ports.EventService
}
