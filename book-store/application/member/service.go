package member

import "github.com/geb/aweproj/book-store/shared"

type (
	Service interface {
		Login(username, password string) (str *string, err error)
		CreateUser() (err error)
	}
	service struct {
		SharedHolder shared.Holder
	}
)

func NewService(holder shared.Holder) (Service, error) {
	return &service{holder}, nil
}
