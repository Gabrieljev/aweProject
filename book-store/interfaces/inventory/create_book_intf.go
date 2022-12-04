package inventory

import (
	"context"
	"github.com/geb/aweproj/book-store/shared/dto/inventory"
)

func (s *service) BulkCreateBook(ctx context.Context, bookReqs []inventory.BookReq) (err error) {
	//books := s.convertToBookDao(bookReqs...)
	return s.ApplicationHolder.Inventory.CreateBook()
}

func (s *service) convertToBookDao(bookReqs ...inventory.BookReq) (books []inventory.BookReq) {
	books = append(books, bookReqs...)
	return
}
