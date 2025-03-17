package events

import domain "github.com/jeissoni/EventLine/internal/domain/entities"

func (r Repository) Guardar(event domain.Event) error {

	_, err := r.Database.Exec("INSERT INTO events (organizer_id, category_id, venue_id, title, slug, description, short_description, start_date, end_date, doors_open_date, is_published, is_featured, is_private, main_image_url, status, ticket_sale_start, ticket_sale_end, event_url, terms_and_conditions, max_tickets_per_person, tags, metadata, seo_title, seo_description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)",
		event.OrganizerID,
		event.CategoryID,
		event.VenueID,
		event.Title,
		event.Slug,
		event.Description,
		event.ShortDescription,
		event.StartDate,
		event.EndDate,
		event.DoorsOpenDate,
		event.IsPublished,
		event.IsFeatured,
		event.IsPrivate,
		event.MainImageURL,
		event.Status,
		event.TicketSaleStart,
		event.TicketSaleEnd,
		event.EventURL,
		event.TermsAndConditions,
		event.MaxTicketsPerPerson,
		event.Tags,
		event.Metadata,
		event.SeoTitle,
		event.SeoDescription,
		event.CreatedAt,
		event.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}
