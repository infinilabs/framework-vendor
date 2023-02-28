package httperr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

type testError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (te testError) Error() string {
	return fmt.Sprintf("%s (%d)", te.Message, te.Code)
}

func TestClient(t *testing.T) {
	transport := Transport{
		Next: roundTripperFunc(func(*http.Request) (*http.Response, error) {
			resp := http.Response{}
			resp.StatusCode = 400
			resp.Body = ioutil.NopCloser(strings.NewReader(`{"message": "cannot frob the grob", "code": 1}`))
			return &resp, nil
		}),
	}

	t.Run("nostruct", func(t *testing.T) {
		client := http.Client{Transport: transport}
		resp, err := client.Get("/foo")
		assert.Nil(t, resp)
		httpErr := err.(*url.Error).Unwrap().(Response)
		assert.Equal(t, "Bad Request", httpErr.Error())
		assert.Equal(t, 400, httpErr.StatusCode)

		body, err := ioutil.ReadAll(httpErr.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"message": "cannot frob the grob", "code": 1}`, string(body))
	})

	t.Run("unmarshal", func(t *testing.T) {
		arg := JSON(testError{})
		arg(&transport)

		client := http.Client{Transport: transport}
		resp, err := client.Get("/foo")
		assert.Nil(t, resp)
		httpErr := err.(*url.Error).Unwrap().(testError)
		assert.Equal(t, "cannot frob the grob (1)", httpErr.Error())
		assert.Equal(t, testError{Message: "cannot frob the grob", Code: 1}, httpErr)
	})

	t.Run("unmarshal bad", func(t *testing.T) {
		transport := Transport{
			Next: roundTripperFunc(func(*http.Request) (*http.Response, error) {
				resp := http.Response{}
				resp.StatusCode = 400
				resp.Body = ioutil.NopCloser(strings.NewReader(`{invalid json`))
				return &resp, nil
			}),
		}

		arg := JSON(testError{})
		arg(&transport)

		client := http.Client{Transport: transport}
		resp, err := client.Get("/foo")
		assert.Nil(t, resp)
		httpErr := err.(*url.Error).Unwrap().(Response)
		assert.Equal(t, "Bad Request", httpErr.Error())
		assert.Equal(t, 400, httpErr.StatusCode)

		body, err := ioutil.ReadAll(httpErr.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{invalid json`, string(body))
	})

	t.Run("success", func(t *testing.T) {
		transport := Transport{
			Next: roundTripperFunc(func(*http.Request) (*http.Response, error) {
				resp := http.Response{}
				resp.StatusCode = 200
				resp.Body = ioutil.NopCloser(strings.NewReader(`{"ok": true}`))
				return &resp, nil
			}),
		}
		client := http.Client{Transport: transport}
		resp, err := client.Get("/foo")
		assert.NoError(t, err)

		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"ok": true}`, string(body))

	})

}
