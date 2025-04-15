package ports

import domain "github.com/jeissoni/EventLine/internal/domain/entities"

type UserService interface {
	Create(user domain.User) error
	GetByID(user_id string) (domain.User, error)
	GetAll() ([]domain.User, error)
	Delete(user_id string) error
	Update(user domain.User) error
	GetByEmail(email string) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	GetByUsernameAndPassword(username string, password string) (domain.User, error)
	GetByEmailAndPassword(email string, password string) (domain.User, error)
}

type UserRepository interface {
	Save(user domain.User) error
	GetByID(user_id string) (domain.User, error)
	GetAll() ([]domain.User, error)
	Delete(user_id string) error
	Update(user domain.User) error
	GetByEmail(email string) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	GetByUsernameAndPassword(username string, password string) (domain.User, error)
	GetByEmailAndPassword(email string, password string) (domain.User, error)
}
