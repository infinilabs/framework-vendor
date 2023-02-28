package httperr

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublic(t *testing.T) {
	testCases := []struct {
		Err        error
		StatusCode int
		Body       string
	}{
		{
			Err: Value{
				Public:     true,
				Status:     "teapot!",
				StatusCode: http.StatusTeapot,
				Err:        fmt.Errorf("cannot frob the grob"),
			},
			StatusCode: http.StatusTeapot,
			Body:       "teapot!\n",
		},
		{
			Err: Value{
				Public: true,
				Status: "teapot!",
				Err:    fmt.Errorf("cannot frob the grob"),
			},
			StatusCode: http.StatusInternalServerError,
			Body:       "teapot!\n",
		},
		{
			Err: Value{
				Public: true,
				Err:    fmt.Errorf("cannot frob the grob"),
			},
			StatusCode: http.StatusInternalServerError,
			Body:       "cannot frob the grob\n",
		},

		// private

		{
			Err: Value{
				Status:     "teapot!",
				StatusCode: http.StatusTeapot,
				Err:        fmt.Errorf("cannot frob the grob"),
			},
			StatusCode: http.StatusTeapot,
			Body:       "teapot!\n",
		},
		{
			Err: Value{
				Status: "teapot!",
				Err:    fmt.Errorf("cannot frob the grob"),
			},
			StatusCode: http.StatusInternalServerError,
			Body:       "teapot!\n",
		},
		{
			Err: Value{
				Err: fmt.Errorf("cannot frob the grob"),
			},
			StatusCode: http.StatusInternalServerError,
			Body:       "Internal Server Error\n",
		},
	}

	for i, testCase := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			h := HandlerFunc(func(http.ResponseWriter, *http.Request) error {
				return testCase.Err
			})

			w := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", "/", nil)
			h.ServeHTTP(w, r)

			assert.Equal(t, testCase.StatusCode, w.Code)
			assert.Equal(t, testCase.Body, string(w.Body.Bytes()))
		})
	}
}
