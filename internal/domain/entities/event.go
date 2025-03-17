package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	EventID             uuid.UUID  `json:"event_id" db:"event_id"`
	OrganizerID         uuid.UUID  `json:"organizer_id" db:"organizer_id"`
	CategoryID          *uuid.UUID `json:"category_id,omitempty" db:"category_id"` // Puntero para permitir valores NULL
	VenueID             *uuid.UUID `json:"venue_id,omitempty" db:"venue_id"`       // Puntero para permitir valores NULL
	Title               string     `json:"title" db:"title"`
	Slug                string     `json:"slug" db:"slug"`
	Description         string     `json:"description" db:"description"`
	ShortDescription    string     `json:"short_description" db:"short_description"`
	StartDate           time.Time  `json:"start_date" db:"start_date"`
	EndDate             time.Time  `json:"end_date" db:"end_date"`
	DoorsOpenDate       *time.Time `json:"doors_open_date,omitempty" db:"doors_open_date"` // Puntero para permitir valores NULL
	IsPublished         bool       `json:"is_published" db:"is_published"`
	IsFeatured          bool       `json:"is_featured" db:"is_featured"`
	IsPrivate           bool       `json:"is_private" db:"is_private"`
	MainImageURL        string     `json:"main_image_url" db:"main_image_url"`
	Status              string     `json:"status" db:"status"`
	TicketSaleStart     *time.Time `json:"ticket_sale_start,omitempty" db:"ticket_sale_start"` // Puntero para permitir valores NULL
	TicketSaleEnd       *time.Time `json:"ticket_sale_end,omitempty" db:"ticket_sale_end"`     // Puntero para permitir valores NULL
	EventURL            string     `json:"event_url" db:"event_url"`
	TermsAndConditions  string     `json:"terms_and_conditions" db:"terms_and_conditions"`
	MaxTicketsPerPerson int        `json:"max_tickets_per_person" db:"max_tickets_per_person"`
	Tags                []string   `json:"tags" db:"tags"`
	Metadata            []byte     `json:"metadata" db:"metadata"` // Usar []byte para JSONB
	SeoTitle            string     `json:"seo_title" db:"seo_title"`
	SeoDescription      string     `json:"seo_description" db:"seo_description"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at"`
}
