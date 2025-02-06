package domain

import (
	"database/sql"
	"time"
)

type Event struct {
	ID            uint           `json:"id"`
	OrganizerID   uint           `json:"organizer_id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Location      string         `json:"location"`
	StartDate     time.Time      `json:"start_date"`
	EndDate       time.Time      `json:"end_date"`
	Status        string         `json:"status"`
	BasePrice     float64        `json:"base_price"`
	TotalCapacity int            `json:"total_capacity"`
	CreatedAt     time.Time      `json:"created_at"`
	ImageUrl      sql.NullString `json:"image_url"`
}
