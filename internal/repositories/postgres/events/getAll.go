package events

import (
	"github.com/jeissoni/EventLine/internal/domain/entities"
)

func (r Repository) GetAll() ([]entities.Event, error) {

	rows, err := r.Database.Query("SELECT * FROM events")

	if err != nil {
		return []entities.Event{}, err
	}

	defer rows.Close()

	var events []entities.Event

	for rows.Next() {
		var event entities.Event
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
			return []entities.Event{}, err
		}

		events = append(events, event)
	}

	return events, nil
}
