package inventory

import (
	"context"
	"github.com/geb/aweproj/ses3/application"
	"github.com/geb/aweproj/ses3/shared"
	"github.com/geb/aweproj/ses3/shared/dto/inventory"
)

type (
	ViewService interface {
		FindBookByPubId(ctx context.Context, pubId int) (books []inventory.BookDto, err error)

		BulkCreateBook(ctx context.Context, bookReqs []inventory.BookReq) (err error)

		UpdateBook(ctx context.Context,id int, bookReqs inventory.BookReq) (err error)

		// - inventory-view-service-end
	}

	service struct {
		SharedHolder      shared.Holder
		ApplicationHolder application.Holder
	}
)

func NewViewService(sh shared.Holder, ah application.Holder) (ViewService, error) {
	return &service{sh, ah}, nil
}
