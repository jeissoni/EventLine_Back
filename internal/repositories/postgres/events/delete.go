package events

func (r Repository) Delete(id int) error {
	_, err := r.Database.Exec("DELETE FROM events WHERE id = $1", id)

	if err != nil {
		return err
	}
	return nil
}
