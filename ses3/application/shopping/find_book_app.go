package shopping

func (s *service) FindBookByPubId(pubId int) (books []Book, err error) {

	ts := s.SharedHolder.DB.Preload("Publisher").Where("pub_id = ?", pubId).Find(&books)
	if ts.Error != nil {
		return nil, err
	}
	return books, nil
}
