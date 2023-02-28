package httperr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddlewareCallsOnErrorOnUnreportedErrors(t *testing.T) {
	var didCallOnError bool
	mw := Middleware{
		OnError: func(w http.ResponseWriter, r *http.Request, err error) error {
			respErr := err.(Response)
			assert.Equal(t, 418, respErr.StatusCode)
			assert.Equal(t, http.Header{"X-Foo": []string{"bar"}}, respErr.Header)

			body, bodyErr := ioutil.ReadAll(respErr.Body)
			assert.NoError(t, bodyErr)
			assert.Equal(t, "response body\n", string(body))

			w.WriteHeader(500)
			fmt.Fprint(w, "ERROR: "+err.Error())
			didCallOnError = true
			return nil
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Foo", "bar")
			w.WriteHeader(418)
			fmt.Fprintln(w, "response body")
		}),
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo", nil)
	mw.ServeHTTP(w, r)

	assert.True(t, didCallOnError, "must call OnError")

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, http.Header{"X-Foo": []string{"bar"}}, w.Header())
	assert.Equal(t, "ERROR: I'm a teapot", string(w.Body.Bytes()))
}

func TestMiddlewareCallsReportsError(t *testing.T) {
	var didCallOnError bool
	mw := Middleware{
		OnError: func(w http.ResponseWriter, r *http.Request, err error) error {
			assert.EqualError(t, err, "cannot frob the grob")
			w.WriteHeader(500)
			fmt.Fprint(w, "ERROR: "+err.Error())
			didCallOnError = true
			return nil
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ReportError(r, fmt.Errorf("cannot frob the grob"))
		}),
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo", nil)
	mw.ServeHTTP(w, r)

	assert.True(t, didCallOnError, "must call OnError")

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "ERROR: cannot frob the grob", string(w.Body.Bytes()))
}

func TestMiddlewareRenderMutatedError(t *testing.T) {
	var didCallOnError bool
	mw := Middleware{
		OnError: func(w http.ResponseWriter, r *http.Request, err error) error {
			assert.EqualError(t, err, "cannot frob the grob")
			didCallOnError = true
			return Public(500, err)
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ReportError(r, fmt.Errorf("cannot frob the grob"))
		}),
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo", nil)
	mw.ServeHTTP(w, r)

	assert.True(t, didCallOnError, "must call OnError")

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "cannot frob the grob\n", string(w.Body.Bytes()))
}

func TestMiddlewareSuccess(t *testing.T) {
	mw := Middleware{
		OnError: func(w http.ResponseWriter, r *http.Request, err error) error {
			assert.Fail(t, "not reached")
			return nil
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Foo", "bar")
			w.WriteHeader(200)
			fmt.Fprintln(w, "response body")
		}),
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo", nil)
	mw.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, http.Header{"X-Foo": []string{"bar"}}, w.Header())
	assert.Equal(t, "response body\n", string(w.Body.Bytes()))
}
