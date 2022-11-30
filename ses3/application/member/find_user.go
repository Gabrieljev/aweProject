package member

import (
	"github.com/pkg/errors"
)

func (s *service) Login(username, password string) (str *string, err error) {
	usr := User{}

	ts := s.SharedHolder.DB.Where("username = ?", username).Find(&usr)
	if ts.Error != nil {
		return nil, ts.Error
	}

	if usr.ID == 0 || !s.comparingHash(usr.Password, password) {
		return nil, errors.New("your username or password mismatch")
	}

	return &usr.Password, nil
}
