package inventory

import (
	"github.com/geb/aweproj/book-store/shared"
	"github.com/geb/aweproj/book-store/shared/dto/inventory"
)

type (
	Service interface {
		CreateBook() error
		FindBookByPubId(pubId int, token string) (bookDto []inventory.BookDto, err error)
		UpdateBook(id int) error
		DeleteBookById(id int) error
		FindPublisher() error
	}
	service struct {
		SharedHolder shared.Holder
	}
)

func NewService(holder shared.Holder) (Service, error) {
	return &service{holder}, nil
}
