package events

import (
	"errors"
	"log"

	custonErrors "github.com/jeissoni/EventLine/internal/domain/custonErrors"
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (s *Service) Update(events domain.Event) error {

	// Check if the event exists

	_, err := s.Repository.GetByID(events.EventID.String())

	if err != nil {

		if errors.Is(err, custonErrors.ErrNotFound) {
			log.Println("No rows in result set")

			appErr := custonErrors.NewDomainError(
				custonErrors.ErrCodeNotFound,
				"Event not found",
			)
			return appErr
		}
		return err
	}

	err = s.Repository.Update(events)
	if err != nil {
		log.Println("Error updating event:", err)
		return err
	}
	log.Println("Event updated successfully")
	return nil

}
