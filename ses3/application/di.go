// Code generated by ntaps. DO NOT EDIT.
package application

import (
	"github.com/geb/aweproj/ses3/application/healthcheck"
	"github.com/geb/aweproj/ses3/application/inventory"
	"github.com/geb/aweproj/ses3/application/member"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {

	// - healthcheck-domain-start
	if err := container.Provide(healthcheck.NewService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck app service")
	}
	// - healthcheck-domain-end

	// - inventory-domain-start
	if err := container.Provide(inventory.NewService); err != nil {
		return errors.Wrap(err, "failed to provide inventory app service")
	}
	// - inventory-domain-end

	// - member-domain-start
	if err := container.Provide(member.NewService); err != nil {
		return errors.Wrap(err, "failed to provide member app service")
	}
	// - member-domain-end

	// application-di-end

	return nil
}
