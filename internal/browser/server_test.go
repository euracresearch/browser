package browser

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gitlab.inf.unibz.it/lter/browser/internal/ql"
)

// testBackend is a implementation of Backend.
// It is meant for tests and does not retrive data from
// a real backend.
type testBackend struct{}

func (tb testBackend) Filter(q ql.Querier) (*Filter, error) {
	return nil, errors.New("not yet implemented")
}

func (tb testBackend) Series(q ql.Querier) ([][]string, error) {
	return nil, errors.New("not yet implemented")
}

func (tb testBackend) Stations(ids ...string) (Stations, error) {
	return nil, errors.New("not yet implemented")
}

type testDecoder struct{}

func (td testDecoder) DecodeAndValidate(r *http.Request) (*Filter, error) {
	return nil, errors.New("not yet implemented")
}

func withTestBackend() Option {
	return func(s *Server) {
		s.db = testBackend{}
	}
}

func withTestBasePath() Option {
	return func(s *Server) {
		s.basePath = "../../static"
	}
}

func WithTestDecoder() Option {
	return func(s *Server) {
		s.decoder = testDecoder{}
	}
}

func TestHandleUpdate(t *testing.T) {
	testCases := []struct {
		method         string
		body           string
		wantStatusCode int
		wantResponse   string
	}{
		{http.MethodGet, "", http.StatusMethodNotAllowed, ""},
		{http.MethodPut, "", http.StatusMethodNotAllowed, ""},
		{http.MethodDelete, "", http.StatusMethodNotAllowed, ""},
		{http.MethodHead, "", http.StatusMethodNotAllowed, ""},
		{http.MethodPatch, "", http.StatusMethodNotAllowed, ""},
		{http.MethodConnect, "", http.StatusMethodNotAllowed, ""},
		{http.MethodOptions, "", http.StatusMethodNotAllowed, ""},
		{http.MethodTrace, "", http.StatusMethodNotAllowed, ""},
		{http.MethodPost, "", http.StatusInternalServerError, ""},
		{http.MethodPost, "{}", http.StatusInternalServerError, ""},
		//{http.MethodPost, "{\"fields\": [\"a\", \"b\", \"c\"]}", http.StatusNotAcceptable, ""},
	}
	for _, tc := range testCases {
		s, err := NewServer(
			withTestBackend(),
			withTestBasePath(),
			WithTestDecoder(),
		)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(tc.method, "/api/v1/filter", strings.NewReader(tc.body))
		w := httptest.NewRecorder()
		s.ServeHTTP(w, r)

		if w.Code != tc.wantStatusCode {
			t.Errorf("got status code %d, want %d", w.Code, tc.wantStatusCode)
		}
	}

}
