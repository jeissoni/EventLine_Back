package events

import domain "github.com/jeissoni/EventLine/internal/domain/entities"

func (r Repository) Guardar(event domain.Event) error {

	_, err := r.Database.Exec("INSERT INTO events (organizer_id, name, description, location, start_date, end_date, status, base_price, total_capacity, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		event.OrganizerID,
		event.Name,
		event.Description,
		event.Location,
		event.StartDate,
		event.EndDate,
		event.Status,
		event.BasePrice,
		event.TotalCapacity,
		event.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}
