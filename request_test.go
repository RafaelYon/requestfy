package requestfy_test

import (
	"net/http"
	"testing"

	"github.com/RafaelYon/requestfy"
)

func TestGet(t *testing.T) {
	t.Run("should concatenate base and specified URL", func(t *testing.T) {
		spy := &spyRequestExecutor{}
		client := requestfy.NewClient(
			requestfy.ConfigRequestExecuter(spy),
			requestfy.ConfigBaseURL("http://some-cool-domain.local"),
		)

		res, err := client.Request().Get("cool/path")
		if err != nil {
			t.Fatalf("expected no error, received '%s'", err)
		}

		if res == nil {
			t.Fatalf("expected non nil *http.Response, received nil")
		}

		if spy.lastRequest == nil {
			t.Fatalf("expected non nil *http.Request, received nil")
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected 200 Status Code, received '%d'", res.StatusCode)
		}

		if spy.lastRequest.Method != http.MethodGet {
			t.Errorf("expected '%s' method, used '%s'", http.MethodGet, spy.lastRequest.Method)
		}

		if expected, used := "http://some-cool-domain.local/cool/path", spy.lastRequest.URL.String(); used != expected {
			t.Errorf("expected '%s' URL, used '%s'", expected, used)
		}
	})
}

func TestHeaders(t *testing.T) {
	t.Run("should add headers to request", func(t *testing.T) {
		cli := requestfy.NewClient()

		testHeaders := map[string][]string{
			"bar":  []string{"foo"},
			"cool": []string{"header", "value"},
		}

		r := cli.Request()

		for k, headers := range testHeaders {
			for _, header := range headers {
				r.SetHeader(k, header)
			}
		}

		for k, wantedHeaders := range testHeaders {
			currentHeaders, ok := r.GetHeaders()[k]

			if !ok {
				t.Errorf("cannot value for %s key header", k)
			}

			for _, wantedHeader := range wantedHeaders {
				if !contains(currentHeaders, wantedHeader) {
					t.Error("headers are not equals")
				}
			}
		}
	})
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type spyRequestExecutor struct {
	lastRequest *http.Request
}

func (s *spyRequestExecutor) Do(req *http.Request) (*http.Response, error) {
	s.lastRequest = req

	return &http.Response{
		StatusCode: http.StatusOK,
	}, nil
}
