package inventory

import (
	"context"
	"github.com/geb/aweproj/book-store/shared/dto/inventory"
)

func (s *service) FindBookByPubId(ctx context.Context, pubId int) (books []inventory.BookDto, err error) {

	//bookRes, err := s.ApplicationHolder.Inventory.FindBookByPubId(pubId)
	//if err != nil {
	//	return nil, err
	//}
	//books = s.convertToBooks(bookRes)
	return
}

func (s *service) convertToBooks(bookDao []inventory.BookDto) (books []inventory.BookDto) {
	books = append(books, bookDao...)
	return
}
