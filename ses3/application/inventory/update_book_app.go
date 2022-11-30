package inventory


func (s *service) UpdateBook(id int, book Book) error {
	tx := s.SharedHolder.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Where("id = ?",id).Updates(book).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}