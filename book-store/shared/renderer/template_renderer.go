package renderer

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"github.com/pkg/errors"
)

type Template struct {
	Templates map[string]*template.Template
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

