package di

import (
	"github.com/geb/aweproj/book-store/shared/renderer"
	"github.com/pkg/errors"
	"html/template"
)

func NewTemplate() *renderer.Template {

	return &renderer.Template{
		Templates: make(map[string]*template.Template),
	}
}

func init() {
	if err := Container.Provide(NewTemplate); err != nil {
		panic(errors.Wrap(err, "failed to provide echo"))
	}
}
