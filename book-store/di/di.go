package di

import(
	"github.com/geb/aweproj/book-store/application"
	"github.com/geb/aweproj/book-store/interfaces"
	"github.com/geb/aweproj/book-store/shared"
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"github.com/geb/aweproj/book-store/infrastructure"
)

var (
	Container *dig.Container = dig.New()
)

func init() {
	if err:= shared.Register(Container); err != nil{
		panic(errors.Wrap(err, "failed to register shared container"))
	}
	if err:= infrastructure.Register(Container); err != nil{
		panic(errors.Wrap(err, "failed to register infrastructure container"))
	}

	if err:= interfaces.Register(Container); err != nil{
		panic(errors.Wrap(err, "failed to register interfaces container"))
	}
	if err:= application.Register(Container); err != nil{
		panic(errors.Wrap(err, "failed to register application container"))
	}
}