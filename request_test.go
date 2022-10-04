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
		assertRequestMethod(t, spy, res, err, http.MethodGet)

		if expected, used := "http://some-cool-domain.local/cool/path", spy.lastRequest.URL.String(); used != expected {
			t.Errorf("expected '%s' URL, used '%s'", expected, used)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("should make a delete http request", func(t *testing.T) {
		spy := &spyRequestExecutor{}

		cli := requestfy.NewClient(
			requestfy.ConfigRequestExecuter(spy),
			requestfy.ConfigBaseURL("http://some-cool-domain.local"),
		)

		res, err := cli.Request().Delete("bar/foo")
		assertRequestMethod(t, spy, res, err, http.MethodDelete)
	})
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
	res *http.Response,
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
