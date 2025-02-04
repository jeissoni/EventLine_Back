package ports

import "github.com/jeissoni/EventLine/internal/domain/entities"

type EventService interface {
	Create(events entities.Event) error
	GetByID(id int) (entities.Event, error)
	GetAll() ([]entities.Event, error)
	Delete(id int) error
}

type EventRepository interface {
	Guardar(events entities.Event) error
	GetByID(id int) (entities.Event, error)
	GetAll() ([]entities.Event, error)
	Delete(id int) error
}
