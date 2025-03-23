package events

func (s *Service) Delete(event_id string) error {
	err := s.Repository.Delete(event_id)
	if err != nil {
		return err
	}
	return nil
}
