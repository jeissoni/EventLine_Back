package user

import (
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (s *Service) Create(user domain.User) error {
	err := s.Repository.Save(user)
	if err != nil {
		return err
	}
	return nil
}
