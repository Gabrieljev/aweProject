package renderer

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"html/template"
	"io"
)

type Template struct {
	Templates    map[string]*template.Template
	FuncTemplate template.FuncMap
}

func (t *Template) Render(w io.Writer, html_name string, data interface{}, c echo.Context) error {
	if tmpl, exist := t.Templates[html_name]; exist {
		return tmpl.Execute(w, data)
	} else {
		return errors.New("There is no " + html_name + " in Template map.")
	}

}

func (tmpl *Template) Add(html_name string, template *template.Template) {
	tmpl.Templates[html_name] = template
}
