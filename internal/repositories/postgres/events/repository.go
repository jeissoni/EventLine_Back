package events

import (
	"database/sql"
)

type Repository struct {
	Database *sql.DB
}
