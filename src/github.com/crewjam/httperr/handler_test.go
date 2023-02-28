package httperr

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	h := HandlerFunc(func(http.ResponseWriter, *http.Request) error {
		return Teapot
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	h.ServeHTTP(w, r)

	assert.Equal(t, 418, w.Code)
	assert.Equal(t, "I'm a teapot\n", string(w.Body.Bytes()))
}
