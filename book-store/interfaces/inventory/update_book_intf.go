package inventory

import (
	"context"
	"github.com/geb/aweproj/book-store/shared/dto/inventory"
)

func (s *service) UpdateBook(ctx context.Context,id int, bookReqs inventory.BookReq) (err error) {
	//books := s.convertToBookDao(bookReqs)
	return s.ApplicationHolder.Inventory.UpdateBook(id)
}
