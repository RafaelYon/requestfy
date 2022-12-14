package requestfy_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/RafaelYon/requestfy"
)

const fakeURL = "http://some-cool-domain.local"
const path = "cool/path"

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

func TestMapHeaders(t *testing.T) {
	t.Run("should set headers by map", func(t *testing.T) {
		cli := requestfy.NewClient()

		testHeaders := http.Header{
			"bar":           []string{"foo", "cool"},
			"header1":       []string{"val1", "val2", "val3"},
			"Authorization": []string{"bearer ..."},
		}

		r := cli.Request().SetHeaders(testHeaders)

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

func TestRequests(t *testing.T) {
	testCases := []struct {
		expectedMethod string
		method         func(*requestfy.Request) func(string, io.Reader) (*requestfy.Response, error)
	}{
		{
			http.MethodGet,
			func(r *requestfy.Request) func(string, io.Reader) (*requestfy.Response, error) {
				return func(s string, body io.Reader) (*requestfy.Response, error) {
					return r.Get(s)
				}
			},
		},
		{
			http.MethodPut,
			func(r *requestfy.Request) func(string, io.Reader) (*requestfy.Response, error) {
				return r.Put
			},
		},
		{
			http.MethodPost,
			func(r *requestfy.Request) func(string, io.Reader) (*requestfy.Response, error) {
				return r.Post
			},
		},
		{
			http.MethodDelete,
			func(r *requestfy.Request) func(string, io.Reader) (*requestfy.Response, error) {
				return func(s string, body io.Reader) (*requestfy.Response, error) {
					return r.Delete(s)
				}
			},
		},
		{
			http.MethodHead,
			func(r *requestfy.Request) func(string, io.Reader) (*requestfy.Response, error) {
				return func(s string, body io.Reader) (*requestfy.Response, error) {
					return r.Head(s)
				}
			},
		},
		{
			http.MethodPatch,
			func(r *requestfy.Request) func(string, io.Reader) (*requestfy.Response, error) {
				return func(s string, body io.Reader) (*requestfy.Response, error) {
					return r.Patch(s)
				}
			},
		},
		{
			http.MethodOptions,
			func(r *requestfy.Request) func(string, io.Reader) (*requestfy.Response, error) {
				return func(s string, body io.Reader) (*requestfy.Response, error) {
					return r.Options(s)
				}
			},
		},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("should make %s http request", test.expectedMethod), func(t *testing.T) {
			spy := &spyRequestExecutor{}
			client := requestfy.NewClient(
				requestfy.ConfigRequestExecuter(spy),
				requestfy.ConfigBaseURL(fakeURL),
			)

			res, err := test.method(client.Request())(path, nil)
			assertRequestMethod(t, spy, res, err, test.expectedMethod)
		})
	}
}

func TestQueryParams(t *testing.T) {
	t.Run("should add params in url", func(t *testing.T) {
		spy := &spyRequestExecutor{}
		client := requestfy.NewClient(
			requestfy.ConfigRequestExecuter(spy),
			requestfy.ConfigBaseURL(fakeURL),
		)

		r := client.Request().
			SetParam("bar", "foo").
			SetParam("query", "string")

		res, err := r.Get(path)

		if err != nil {
			t.Error(err)
		}

		urlExpected := fmt.Sprintf("%s/%s?bar=foo&query=string", fakeURL, path)
		url := res.Response.Request.URL.String()

		if url != urlExpected {
			t.Errorf("url expected: %s but got %s", urlExpected, url)
		}
	})
}

type spyRequestExecutor struct {
	lastRequest *http.Request
}

func (s *spyRequestExecutor) Do(req *http.Request) (*http.Response, error) {
	s.lastRequest = req

	return &http.Response{
		StatusCode: http.StatusOK,
		Request:    req,
	}, nil
}

func assertRequestMethod(
	t *testing.T,
	spy *spyRequestExecutor,
	res *requestfy.Response,
	err error,
	expectedMethod string,
) {
	t.Helper()

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

	if spy.lastRequest.Method != expectedMethod {
		t.Errorf("expected '%s' method, used '%s'", expectedMethod, spy.lastRequest.Method)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
