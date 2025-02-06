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
			&event.ID,
			&event.OrganizerID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.StartDate,
			&event.EndDate,
			&event.Status,
			&event.BasePrice,
			&event.TotalCapacity,
			&event.CreatedAt,
			&event.ImageUrl,
		)

		if err != nil {
			return []domain.Event{}, err
		}

		events = append(events, event)
	}

	return events, nil
}
