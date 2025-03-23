package events

func (r Repository) Delete(event_id string) error {
	_, err := r.Database.Exec("DELETE FROM events WHERE event_id = $1", event_id)

	if err != nil {
		return err
	}
	return nil
}
