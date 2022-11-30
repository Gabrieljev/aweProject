package member

import "github.com/geb/aweproj/ses3/shared"

type (
	Service interface {
		Login(username, password string) (str *string, err error)
		CreateUser(usr User) (err error)
	}
	service struct {
		SharedHolder shared.Holder
	}
)

func NewService(holder shared.Holder) (Service, error) {
	return &service{holder}, nil
}
