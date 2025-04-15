package user

import (
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (r Repository) Guardar(user domain.User) error {
	_, err := r.Database.Exec(
		`INSERT INTO users (
            user_id, email, password_hash, first_name, last_name, phone, date_of_birth, 
            profile_picture_url, is_verified, verification_token, reset_password_token, 
            reset_password_expires, last_login, role, created_at, updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
        )`,
		user.UserID,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.DateOfBirth,
		user.ProfilePictureURL,
		user.IsVerified,
		user.VerificationToken,
		user.ResetPasswordToken,
		user.ResetPasswordExpires,
		user.LastLogin,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
