package shopping

import (
	"context"
	shopping_dao "github.com/geb/aweproj/ses3/application/shopping"
	"github.com/geb/aweproj/ses3/shared/dto/shopping"
)

func (s *service) BulkCreateBook(ctx context.Context, bookReqs []shopping.CreateBookReq) (err error) {
	books := s.convertToBookDao(bookReqs)
	return s.ApplicationHolder.Shopping.CreateBook(books)
}

func (s *service) convertToBookDao(bookReqs []shopping.CreateBookReq) (books []shopping_dao.Book) {
	books = make([]shopping_dao.Book, 0)
	for _, v := range bookReqs {
		books = append(books, shopping_dao.Book{
			Title:  v.Title,
			Author: v.Author,
			PubId:  v.PubId,
		})
	}
	return
}
