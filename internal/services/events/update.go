package events

import (
	"log"

	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (s *Service) Update(events domain.Event) error {

	err := s.Repository.Update(events)
	if err != nil {
		log.Println("Error updating event:", err)
		return err
	}
	log.Println("Event updated successfully")
	return nil

}
