package inventory

func (s *service) DeleteBookById(id int) error {
	tx := s.SharedHolder.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Table("book").Where("id = ?", id).Update("is_deleted", 1).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
