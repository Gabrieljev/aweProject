package inventory

import (
	"context"
	"github.com/geb/aweproj/ses3/shared/dto/inventory"
)

func (s *service) UpdateBook(ctx context.Context,id int, bookReqs inventory.BookReq) (err error) {
	books := s.convertToBookDao(bookReqs)
	return s.ApplicationHolder.Inventory.UpdateBook(id, books[0])
}
