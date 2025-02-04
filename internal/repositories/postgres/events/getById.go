package events

import (
	"database/sql"
	"errors"

	"github.com/jeissoni/EventLine/internal/domain/entities"
)

var ErrNotFound = errors.New("event not found")

func (r Repository) GetByID(id int) (entities.Event, error) {
	var event entities.Event
	err := r.Database.QueryRow("SELECT * FROM events WHERE id = $1", id).Scan(
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

	if err != sql.ErrNoRows {
		return entities.Event{}, ErrNotFound
	}

	if err != nil {
		return entities.Event{}, err
	}

	return event, nil
}
