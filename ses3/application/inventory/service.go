package inventory

import "github.com/geb/aweproj/ses3/shared"

type (
	Service interface {
		CreateBook(book []Book) error
		FindBookByPubId(pubId int) (books []Book, err error)
		UpdateBook(id int, book Book) error
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
