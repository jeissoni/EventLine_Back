package events

import (
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (r Repository) GetAll() ([]domain.Event, error) {

	rows, err := r.Database.Query("SELECT * FROM events")

	if err != nil {
		return []domain.Event{}, err
	}

	var events []domain.Event

	for rows.Next() {
		var event domain.Event
		err := rows.Scan(
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
			return []domain.Event{}, err
		}

		events = append(events, event)
	}

	return events, nil
}
