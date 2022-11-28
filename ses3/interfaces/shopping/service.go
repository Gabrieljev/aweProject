package shopping

import (
	"context"
	"github.com/geb/aweproj/ses3/application"
	"github.com/geb/aweproj/ses3/shared"
	"github.com/geb/aweproj/ses3/shared/dto/shopping"
)

type (
	ViewService interface {

		// - healthcheck-check-start
		FindBookByPubId(ctx context.Context,pubId int) (books []shopping.BookDto, err error)
		// - healthcheck-check-end

		// - healthcheck-view-service-end
	}

	service struct {
		SharedHolder      shared.Holder
		ApplicationHolder application.Holder
	}
)

func NewViewService(sh shared.Holder, ah application.Holder) (ViewService, error) {
	return &service{sh, ah}, nil
}