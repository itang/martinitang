package martinitang

import (
	"html/template"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func MakeCommonRender(layout string, funcMaps ...template.FuncMap) martini.Handler {
	return render.Renderer(render.Options{
		Delims:     render.Delims{"{[{", "}]}"},
		Directory:  ".",                        // specify what path to load the templates from
		Layout:     layout,                     // specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates
		Funcs:      funcMaps,                   // Specify helper function maps for templates to access.
	})
}

func DefaultRender() martini.Handler {
	return render.Renderer(render.Options{})
}
