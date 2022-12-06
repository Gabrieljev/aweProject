// Code generated by ntaps. DO NOT EDIT.
package application

import (
	"github.com/geb/aweproj/book-store/application/inventory"
	"github.com/geb/aweproj/book-store/application/member"
	"go.uber.org/dig"

	"github.com/geb/aweproj/book-store/application/healthcheck"

)

type (
	Holder struct {
		dig.In

		// - healthcheck-domain-start
		HealthCheckService healthcheck.Service
		// - healthcheck-domain-end

		// - inventory-domain-start
		Inventory inventory.Service
		// - inventory-domain-end

		// - member-domain-start
		Member member.Service
		// - member-domain-end


		// - application-holder-end
	}
)