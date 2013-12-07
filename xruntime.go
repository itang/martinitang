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

func xRuntimeHandler(w http.ResponseWriter, r *http.Request, c martini.Context) {
	xw := xRuntimeResponseWriter{time.Now(), martini.NewResponseWriter(w)}
	xw.Before(func(w martini.ResponseWriter) {
		xw.tryAddXRuntimeHeader()
	})
	c.MapTo(xw, (*http.ResponseWriter)(nil))

	c.Next()
}

type xRuntimeResponseWriter struct {
	start time.Time
	martini.ResponseWriter
}

// use martini.ResponseWriter#Before
/*
func (xw xRuntimeResponseWriter) WriteHeader(s int) {
	xw.tryAddXRuntimeHeader()
	xw.ResponseWriter.WriteHeader(s)
}

func (xw xRuntimeResponseWriter) Write(p []byte) (int, error) {
	xw.tryAddXRuntimeHeader()
	return xw.ResponseWriter.Write(p)
}
*/

func (xw xRuntimeResponseWriter) tryAddXRuntimeHeader() {
	if len(xw.Header().Get(HeaderXRuntime)) == 0 {
		xw.Header().Add(HeaderXRuntime, fmt.Sprintf("%v", time.Since(xw.start)))
	}
}
