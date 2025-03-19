package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	EventID             uuid.UUID  `json:"event_id" db:"event_id"`
	OrganizerID         uuid.UUID  `json:"organizer_id" db:"organizer_id"`
	CategoryID          *uuid.UUID `json:"category_id,omitempty" db:"category_id"` // Puntero para permitir valores NULL
	Title               string     `json:"title" db:"title"`
	Description         string     `json:"description" db:"description"`
	ShortDescription    string     `json:"short_description" db:"short_description"`
	StartDate           time.Time  `json:"start_date" db:"start_date"`
	EndDate             time.Time  `json:"end_date" db:"end_date"`
	IsPublished         bool       `json:"is_published" db:"is_published"`
	IsFeatured          bool       `json:"is_featured" db:"is_featured"`
	IsPrivate           bool       `json:"is_private" db:"is_private"`
	MainImageURL        string     `json:"main_image_url" db:"main_image_url"`
	Status              string     `json:"status" db:"status"`
	TicketSaleStart     *time.Time `json:"ticket_sale_start,omitempty" db:"ticket_sale_start"` // Puntero para permitir valores NULL
	TicketSaleEnd       *time.Time `json:"ticket_sale_end,omitempty" db:"ticket_sale_end"`     // Puntero para permitir valores NULL
	MaxTicketsPerPerson int        `json:"max_tickets_per_person" db:"max_tickets_per_person"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at"`
}
