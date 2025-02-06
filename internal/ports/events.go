package ports

import domain "github.com/jeissoni/EventLine/internal/domain/entities"

type EventService interface {
	Create(events domain.Event) error
	GetByID(id int) (domain.Event, error)
	GetAll() ([]domain.Event, error)
	Delete(id int) error
}

type EventRepository interface {
	Guardar(events domain.Event) error
	GetByID(id int) (domain.Event, error)
	GetAll() ([]domain.Event, error)
	Delete(id int) error
}
