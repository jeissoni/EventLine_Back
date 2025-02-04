package events

import (
	"github.com/jeissoni/EventLine/internal/domain/entities"
)

func (s *Service) GetByID(id int) (entities.Event, error) {
	event, err := s.Repository.GetByID(id)
	if err != nil {

		return entities.Event{}, err
	}
	return event, nil
}
