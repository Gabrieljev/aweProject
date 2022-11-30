package member

func (s *service) CreateUser(usr User) (err error) {
	usr.Password = s.encrypt( usr.Password)
	tx := s.SharedHolder.DB.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(&usr).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
