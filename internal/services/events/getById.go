package events

import (
	"errors"
	"log"

	custonErrors "github.com/jeissoni/EventLine/internal/domain/custonErrors"
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (s *Service) GetByID(id int) (domain.Event, error) {
	event, err := s.Repository.GetByID(id)
	if err != nil {

		if errors.Is(err, custonErrors.ErrNotFound) {
			log.Println("No rows in result set")

			appErr := custonErrors.NewDomainError(
				custonErrors.ErrCodeNotFound,
				"Event not found",
			)

			return domain.Event{}, appErr
		}

		return domain.Event{}, err
	}
	return event, nil
}
