package inventory

func (s *service) FindBookByPubId(pubId int) (books []Book, err error) {

	ts := s.SharedHolder.DB.Preload("Publisher").Where("pub_id = ? and is_deleted = ?", pubId, 0).Find(&books)
	if ts.Error != nil {
		return nil, err
	}
	return books, nil
}
