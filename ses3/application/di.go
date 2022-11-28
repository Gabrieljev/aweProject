// Code generated by ntaps. DO NOT EDIT.
package application

import (
	"github.com/geb/aweproj/ses3/application/shopping"
	"github.com/pkg/errors"

	"github.com/geb/aweproj/ses3/application/healthcheck"


	"go.uber.org/dig"
)

func Register(container *dig.Container) error {


	// - healthcheck-domain-start
	if err := container.Provide(healthcheck.NewService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck app service")
	}
	// - healthcheck-domain-end

	// - shopping-domain-start
	if err := container.Provide(shopping.NewService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck app service")
	}
	// - shopping-domain-end

	// application-di-end

	return nil
}
