package events

import (
	"database/sql"
	"fmt"

	custonErrors "github.com/jeissoni/EventLine/internal/domain/custonErrors"
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (r Repository) GetByID(event_id string) (domain.Event, error) {
	var event domain.Event
	err := r.Database.QueryRow("SELECT * FROM events WHERE event_id = $1", event_id).Scan(
		&event.EventID,
		&event.OrganizerID,
		&event.CategoryID,
		&event.Title,
		&event.Description,
		&event.ShortDescription,
		&event.StartDate,
		&event.EndDate,
		&event.IsPublished,
		&event.IsFeatured,
		&event.IsPrivate,
		&event.MainImageURL,
		&event.Status,
		&event.TicketSaleStart,
		&event.TicketSaleEnd,
		&event.MaxTicketsPerPerson,
		&event.CreatedAt,
		&event.UpdatedAt,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return domain.Event{},
				fmt.Errorf("%w: %s", custonErrors.ErrNotFound, err.Error())
		}
		return domain.Event{}, err
	}

	return event, nil
}
