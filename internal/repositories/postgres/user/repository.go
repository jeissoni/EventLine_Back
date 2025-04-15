package user

import (
	"database/sql"
)

type Repository struct {
	Database *sql.DB
}
