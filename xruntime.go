package martinitang

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
)

const HeaderXRuntime = "X-Runtime"

func XRuntime() martini.Handler {
	return xRuntimeHandler
}

func xRuntimeHandler(w http.ResponseWriter, c martini.Context) {
	start := time.Now()
	xw := martini.NewResponseWriter(w)
	xw.Before(func(w martini.ResponseWriter) {
		if len(xw.Header().Get(HeaderXRuntime)) == 0 {
			xw.Header().Add(HeaderXRuntime, fmt.Sprintf("%v", time.Since(start)))
		}
	})
	c.MapTo(xw, (*http.ResponseWriter)(nil))

	c.Next()
}
