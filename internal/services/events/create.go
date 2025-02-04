package events

import (
	"github.com/jeissoni/EventLine/internal/domain/entities"
)

func (s *Service) Create(events entities.Event) error {
	err := s.Repository.Guardar(events)
	if err != nil {
		return err
	}
	return nil
}
