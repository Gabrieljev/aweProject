package inventory

import (
	"context"
	inventory_dao "github.com/geb/aweproj/ses3/application/inventory"
	"github.com/geb/aweproj/ses3/shared/dto/inventory"
)

func (s *service) FindBookByPubId(ctx context.Context, pubId int) (books []inventory.BookDto, err error) {

	bookRes, err := s.ApplicationHolder.Inventory.FindBookByPubId(pubId)
	if err != nil {
		return nil, err
	}
	books = s.convertToBooks(bookRes)
	return
}

func (s *service) convertToBooks(bookDao []inventory_dao.Book) (books []inventory.BookDto) {
	books = make([]inventory.BookDto, 0)
	for _, v := range bookDao {
		books = append(books, inventory.BookDto{
			ID:     int(v.ID),
			Title:  v.Title,
			Author: v.Author,
			Publisher: inventory.PublisherDto{
				ID:        int(v.Publisher.ID),
				Name:      v.Publisher.Name,
				CountryId: int(v.Publisher.CountryId),
			},
		})
	}
	return
}
