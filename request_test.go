package requestfy_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/RafaelYon/requestfy"
)

const fakeURL = "http://some-cool-domain.local"
const path = "cool/path"

func TestRequests(t *testing.T) {
	testCases := []struct {
		expectedMethod string
		method         func(*requestfy.Request) func(string) (*http.Response, error)
	}{
		{
			http.MethodGet,
			func(r *requestfy.Request) func(string) (*http.Response, error) {
				return r.Get
			},
		},
		{
			http.MethodDelete,
			func(r *requestfy.Request) func(string) (*http.Response, error) {
				return r.Delete
			},
		},
		{
			http.MethodHead,
			func(r *requestfy.Request) func(string) (*http.Response, error) {
				return r.Head
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

			res, err := test.method(client.Request())(path)
			assertRequestMethod(t, spy, res, err, test.expectedMethod)
		})
	}
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
