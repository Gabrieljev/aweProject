// Code generated by ntaps. DO NOT EDIT.
package di

import (
	"github.com/pkg/errors"

	echo "github.com/labstack/echo/v4"

)

func NewEcho() (*echo.Echo, error) {
	e := echo.New()
	return e, nil
}

func init() {
	if err := Container.Provide(NewEcho); err != nil {
		panic(errors.Wrap(err, "failed to provide echo"))
	}
}