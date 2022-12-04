package inventory

import (
	"context"
)

func (s *service) DeleteBook(ctx context.Context,id int) (err error) {
	return s.ApplicationHolder.Inventory.DeleteBookById(id)
}