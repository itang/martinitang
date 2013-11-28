package martinitang

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/codegangsta/martini"
)

func TestXRuntime(t *testing.T) {
	m, r := writenHandler()
	m.ServeHTTP(r, (*http.Request)(nil))
	v := r.Header().Get("X-Runtime")
	println(v)
	assert.NotEqual(t, "", v, "X-Runtime header shoud be set by XRuntime-Middleware when write data")

	m1, r := emptyHandler()
	m1.ServeHTTP(r, (*http.Request)(nil))
	v = r.Header().Get("X-Runtime")
	assert.Equal(t, "", v, "X-Runtime header shoud not be set by XRuntime-Middleware when no write data")
}

func mr(h martini.Handler) (http.Handler, http.ResponseWriter) {
	m := martini.New()
	m.Use(XRuntime())
	m.Use(h)
	return m, httptest.NewRecorder()
}

func writenHandler() (http.Handler, http.ResponseWriter) {
	return mr(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Millisecond * 100)
		res.Write([]byte("hello"))
	})
}

func emptyHandler() (http.Handler, http.ResponseWriter) {
	return mr(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Millisecond * 100)
	})
}
