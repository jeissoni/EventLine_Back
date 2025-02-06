package events

import (
	"database/sql"
	"fmt"

	custonErrors "github.com/jeissoni/EventLine/internal/domain/custonErrors"
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (r Repository) GetByID(id int) (domain.Event, error) {
	var event domain.Event
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

	if err != nil {

		if err == sql.ErrNoRows {
			return domain.Event{},
				fmt.Errorf("%w: %s", custonErrors.ErrNotFound, err.Error())
		}

		return domain.Event{}, err
	}

	return event, nil
}
