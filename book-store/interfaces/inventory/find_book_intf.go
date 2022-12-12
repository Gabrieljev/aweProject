package inventory

import (
	"context"
	"github.com/geb/aweproj/book-store/shared/dto/inventory"
)

func (s *service) FindBookByPubId(ctx context.Context, pubId int, token string) (books []inventory.BookDto, err error) {

	bookRes, err := s.ApplicationHolder.Inventory.FindBookByPubId(pubId, token)
	if err != nil {
		return nil, err
	}
	//books = s.convertToBooks(bookRes)
	return bookRes, nil
}

func (s *service) convertToBooks(bookDao []inventory.BookDto) (books []inventory.BookDto) {
	books = append(books, bookDao...)
	return
}
