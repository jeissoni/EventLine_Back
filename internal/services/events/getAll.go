package events

import domain "github.com/jeissoni/EventLine/internal/domain/entities"

func (s *Service) GetAll() ([]domain.Event, error) {
	event, err := s.Repository.GetAll()
	if err != nil {

		return []domain.Event{}, err
	}
	return event, nil
}
