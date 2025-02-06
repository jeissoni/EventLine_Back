package events

import (
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (s *Service) Create(events domain.Event) error {
	err := s.Repository.Guardar(events)
	if err != nil {
		return err
	}
	return nil
}
