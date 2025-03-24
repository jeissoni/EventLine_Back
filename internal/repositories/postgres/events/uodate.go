package events

import domain "github.com/jeissoni/EventLine/internal/domain/entities"

func (r Repository) Update(event domain.Event) error {

	_, err := r.Database.Exec(
		"UPDATE events SET organizer_id = $1, category_id = $2, title = $3, description = $4, short_description = $5, start_date = $6, end_date = $7, is_published = $8, is_featured = $9, is_private = $10, main_image_url = $11, status = $12, ticket_sale_start = $13, ticket_sale_end = $14, max_tickets_per_person = $15 WHERE event_id = $16",
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
		event.EventID,
	)

	if err != nil {
		return err
	}

	return nil
}
