package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID               uuid.UUID  `json:"user_id" db:"user_id"`
	Email                string     `json:"email" db:"email"`
	PasswordHash         string     `json:"password_hash" db:"password_hash"`
	FirstName            string     `json:"first_name" db:"first_name"`
	LastName             string     `json:"last_name" db:"last_name"`
	Phone                *string    `json:"phone,omitempty" db:"phone"`                             // Puntero para permitir valores NULL
	DateOfBirth          *time.Time `json:"date_of_birth,omitempty" db:"date_of_birth"`             // Puntero para permitir valores NULL
	ProfilePictureURL    *string    `json:"profile_picture_url,omitempty" db:"profile_picture_url"` // Puntero para permitir valores NULL
	IsVerified           bool       `json:"is_verified" db:"is_verified"`
	VerificationToken    *uuid.UUID `json:"verification_token,omitempty" db:"verification_token"`         // Puntero para permitir valores NULL
	ResetPasswordToken   *uuid.UUID `json:"reset_password_token,omitempty" db:"reset_password_token"`     // Puntero para permitir valores NULL
	ResetPasswordExpires *time.Time `json:"reset_password_expires,omitempty" db:"reset_password_expires"` // Puntero para permitir valores NULL
	LastLogin            *time.Time `json:"last_login,omitempty" db:"last_login"`                         // Puntero para permitir valores NULL
	Role                 string     `json:"role" db:"role"`
	CreatedAt            time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at" db:"updated_at"`
}
