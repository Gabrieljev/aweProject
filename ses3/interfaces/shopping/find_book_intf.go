package shopping

import (
	"context"
	shopping_dao "github.com/geb/aweproj/ses3/application/shopping"
	"github.com/geb/aweproj/ses3/shared/dto/shopping"
)

func (s *service) FindBookByPubId(ctx context.Context, pubId int) (books []shopping.BookDto, err error) {

	bookRes, err := s.ApplicationHolder.Shopping.FindBookByPubId(pubId)
	if err != nil {
		return nil, err
	}
	books = s.convertToBooks(bookRes)
	return
}

func (s *service) convertToBooks(bookDao []shopping_dao.Book) (books []shopping.BookDto) {
	books = make([]shopping.BookDto, 0)
	for _, v := range bookDao {
		books = append(books, shopping.BookDto{
			ID:     int(v.ID),
			Title:  v.Title,
			Author: v.Author,
			Publisher: shopping.PublisherDto{
				ID:        int(v.Publisher.ID),
				Name:      v.Publisher.Name,
				CountryId: int(v.Publisher.CountryId),
			},
		})
	}
	return
}
