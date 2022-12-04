// Code generated by ntaps. DO NOT EDIT.
package di

import (
	"github.com/geb/aweproj/book-store/shared"
	"github.com/pkg/errors"
)

func NewValidator() (shared.Validator, error) {
	return shared.NewValidator(), nil
}

func init() {
	if err := Container.Provide(NewValidator); err != nil {
		panic(errors.Wrap(err, "failed to provide validator"))
	}
}
