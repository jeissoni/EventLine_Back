package events

import (
	"github.com/jeissoni/EventLine/internal/domain/entities"
)

func (s *Service) GetAll() ([]entities.Event, error) {
	event, err := s.Repository.GetAll()
	if err != nil {

		return []entities.Event{}, err
	}
	return event, nil
}
