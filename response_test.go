package requestfy_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/RafaelYon/requestfy"
)

func TestResponseJsonDecode(t *testing.T) {
	stub := &requestExecuterStub{
		response: &http.Response{
			Body: newStringReaderCloserDummy(
				`{"people":"https://swapi.dev/api/people/","planets":"https://swapi.dev/api/planets/","films":"https://swapi.dev/api/films/","species":"https://swapi.dev/api/species/","vehicles":"https://swapi.dev/api/vehicles/","starships":"https://swapi.dev/api/starships/"}`,
			),
		},
	}

	testCases := []struct {
		name    string
		decoder requestfy.NewDecoder
	}{
		{
			name:    "http.Json",
			decoder: requestfy.StdJsonDecoder,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			client := requestfy.NewClient(
				requestfy.ConfigRequestExecuter(stub),
				requestfy.ConfigJsonDecoder(testCase.decoder),
			)

			res, err := client.Request().Get("https://swapi.dev/api/")

			assertNoError(t, err)
			if res == nil {
				t.Fatal("expected non nil response, received nil")
			}

			var endpoints map[string]string
			assertNoError(t, res.Json(&endpoints))

			if total := len(endpoints); total != 6 {
				t.Errorf("expected 6 json decoded endpoints, received '%d'", total)
			}

			expectedEndpoints := map[string]string{
				"people":    "https://swapi.dev/api/people/",
				"planets":   "https://swapi.dev/api/planets/",
				"films":     "https://swapi.dev/api/films/",
				"species":   "https://swapi.dev/api/species/",
				"vehicles":  "https://swapi.dev/api/vehicles/",
				"starships": "https://swapi.dev/api/starships/",
			}
			for expectedKey, expectedVal := range expectedEndpoints {
				if val := endpoints[expectedKey]; val != expectedVal {
					t.Errorf("expected '%s' endpoint in the key '%s', received '%s'", expectedVal, expectedKey, val)
				}
			}
		})
	}
}

type requestExecuterStub struct {
	response *http.Response
	err      error
}

func (r *requestExecuterStub) Do(*http.Request) (*http.Response, error) {
	return r.response, r.err
}

type stringReaderCloserDummy struct {
	*strings.Reader
}

func (*stringReaderCloserDummy) Close() error {
	return nil
}

func newStringReaderCloserDummy(s string) *stringReaderCloserDummy {
	return &stringReaderCloserDummy{
		Reader: strings.NewReader(s),
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected nil error, receive '%s'", err)
	}
}
