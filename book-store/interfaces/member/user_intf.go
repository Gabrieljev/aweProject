package member

import (
	"context"
)

func (s *service) Login(ctx context.Context, username, password string) (token *string, err error) {
	encryptedPwd, err := s.ApplicationHolder.Member.Login(username, password)
	if err != nil {
		return nil, err
	}

	token = encryptedPwd
	return
}

func (s *service) CreateUser(ctx context.Context, username, password string) (err error) {
	return s.ApplicationHolder.Member.CreateUser()
}
