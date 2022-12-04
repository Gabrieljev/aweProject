// Code generated by ntaps. DO NOT EDIT.
package interfaces

import (
	"github.com/geb/aweproj/book-store/interfaces/inventory"
	"github.com/geb/aweproj/book-store/interfaces/member"
	"github.com/pkg/errors"

	"github.com/geb/aweproj/book-store/interfaces/healthcheck"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {

	// - healthcheck-view-service-start
	if err := container.Provide(healthcheck.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck app service")
	}
	// - healthcheck-view-service-end

	// - inventory-domain-start
	if err := container.Provide(inventory.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide inventory app service")
	}
	// - inventory-domain-end

	// - member-domain-start
	if err := container.Provide(member.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide member app service")
	}
	// - member-domain-end

	return nil
}
