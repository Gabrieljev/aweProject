package member

import (
	"context"
	"github.com/geb/aweproj/ses3/application/member"
	"github.com/geb/aweproj/ses3/shared"
)

func (s *service) Login(ctx context.Context, username, password string) (token *string, err error) {
	encryptedPwd, err := s.ApplicationHolder.Member.Login(username, password)
	if err != nil {
		return nil, err
	}
	tk, err := shared.GetToken(username, *encryptedPwd)
	token = &tk
	return
}

func (s *service) CreateUser(ctx context.Context, username, password string) (err error) {
	return s.ApplicationHolder.Member.CreateUser(member.User{Username: username, Password: password})
}
