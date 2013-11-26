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
	c.MapTo(xw, (*http.ResponseWriter)(nil))

	c.Next()
}

type xRuntimeResponseWriter struct {
	start time.Time
	martini.ResponseWriter
}

func (xw xRuntimeResponseWriter) Write(p []byte) (int, error) {
	if len(xw.Header().Get(HeaderXRuntime)) == 0 {
		xw.Header().Add(HeaderXRuntime, fmt.Sprintf("%v", time.Now().Sub(xw.start)))
	}

	return xw.ResponseWriter.Write(p)
}
