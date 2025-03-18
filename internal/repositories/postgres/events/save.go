package events

import domain "github.com/jeissoni/EventLine/internal/domain/entities"

func (r Repository) Guardar(event domain.Event) error {

	_, err := r.Database.Exec(
		"INSERT INTO events (organizer_id, category_id, title, description, short_description, start_date, end_date, is_published, is_featured, is_private, main_image_url, status, ticket_sale_start, ticket_sale_end, max_tickets_per_person) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)",
		event.OrganizerID,
		event.CategoryID,
		event.Title,
		event.Description,
		event.ShortDescription,
		event.StartDate,
		event.EndDate,
		event.IsPublished,
		event.IsFeatured,
		event.IsPrivate,
		event.MainImageURL,
		event.Status,
		event.TicketSaleStart,
		event.TicketSaleEnd,
		event.MaxTicketsPerPerson,
	)

	if err != nil {
		return err
	}

	return nil
}
