package inventory

import (
	"context"
	inventory_dao "github.com/geb/aweproj/ses3/application/inventory"
	"github.com/geb/aweproj/ses3/shared/dto/inventory"
)

func (s *service) BulkCreateBook(ctx context.Context, bookReqs []inventory.BookReq) (err error) {
	books := s.convertToBookDao(bookReqs...)
	return s.ApplicationHolder.Inventory.CreateBook(books)
}

func (s *service) convertToBookDao(bookReqs ...inventory.BookReq) (books []inventory_dao.Book) {
	books = make([]inventory_dao.Book, 0)
	for _, v := range bookReqs {
		books = append(books, inventory_dao.Book{
			Title:  v.Title,
			Author: v.Author,
			PubId:  v.PubId,
		})
	}
	return
}
