package member

import (
	"context"
	"github.com/geb/aweproj/book-store/application"
	"github.com/geb/aweproj/book-store/shared"
)

type (
	ViewService interface {
		Login(ctx context.Context, username, password string) (str *string, err error)
		CreateUser(ctx context.Context, username, password string) (err error)

		// - member-view-service-end
	}

	service struct {
		SharedHolder      shared.Holder
		ApplicationHolder application.Holder
	}
)

func NewViewService(sh shared.Holder, ah application.Holder) (ViewService, error) {
	return &service{sh, ah}, nil
}
