package shopping

import "github.com/geb/aweproj/ses3/shared"

type (
	Service interface {
		InsertBook() error
		FindBookByPubId(pubId int) (books []Book, err error)
		UpdateBook() error
		DeleteBook() error
		FindPublisher() error
	}
	service struct {
		SharedHolder shared.Holder
	}
)

func NewService(holder shared.Holder) (Service, error) {
	return &service{holder}, nil
}
