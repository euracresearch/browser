// Copyright 2020 Eurac Research. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHealthz(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "default")
	})

	t.Run("ok", func(t *testing.T) {
		fn := func() error {
			return nil
		}

		mw := Healthz("token", fn)
		ts := httptest.NewServer(mw(handler))
		defer ts.Close()

		resp, err := ts.Client().Get(MustParseURL(t, ts.URL, "healthz?ht=token"))
		if err != nil {
			t.Fatal(err)
		}

		if got, want := resp.StatusCode, http.StatusOK; got != want {
			t.Fatalf("want status code %d, got %d", want, got)
		}

		if got, want := MustReadBody(t, resp), "ok"; got != want {
			t.Fatalf("want body %q, got %q", want, got)
		}
	})

	t.Run("wrongEndpoint", func(t *testing.T) {
		fn := func() error {
			return nil
		}

		mw := Healthz("token", fn)
		ts := httptest.NewServer(mw(handler))
		defer ts.Close()

		resp, err := ts.Client().Get(MustParseURL(t, ts.URL, "something"))
		if err != nil {
			t.Fatal(err)
		}

		if got, want := resp.StatusCode, http.StatusOK; got != want {
			t.Fatalf("want status code %d, got %d", want, got)
		}

		if got, want := MustReadBody(t, resp), "default"; got != want {
			t.Fatalf("want body %q, got %q", want, got)
		}
	})

	t.Run("unauthorizedEmptyToken", func(t *testing.T) {
		fn := func() error {
			return nil
		}

		mw := Healthz("token", fn)
		ts := httptest.NewServer(mw(handler))
		defer ts.Close()

		resp, err := ts.Client().Get(MustParseURL(t, ts.URL, "healthz?ht="))
		if err != nil {
			t.Fatal(err)
		}

		if got, want := resp.StatusCode, http.StatusUnauthorized; got != want {
			t.Fatalf("want status code %d, got %d", want, got)
		}

		if got, want := MustReadBody(t, resp), "unauthorized\n"; got != want {
			t.Fatalf("want body %q, got %q", want, got)
		}
	})

	t.Run("unauthorizedMissingToken", func(t *testing.T) {
		fn := func() error {
			return nil
		}

		mw := Healthz("token", fn)
		ts := httptest.NewServer(mw(handler))
		defer ts.Close()

		resp, err := ts.Client().Get(MustParseURL(t, ts.URL, "healthz"))
		if err != nil {
			t.Fatal(err)
		}

		if got, want := resp.StatusCode, http.StatusUnauthorized; got != want {
			t.Fatalf("want status code %d, got %d", want, got)
		}

		if got, want := MustReadBody(t, resp), "unauthorized\n"; got != want {
			t.Fatalf("want body %q, got %q", want, got)
		}
	})

	t.Run("unauthorizedWrongToken", func(t *testing.T) {
		fn := func() error {
			return nil
		}

		mw := Healthz("token", fn)
		ts := httptest.NewServer(mw(handler))
		defer ts.Close()

		resp, err := ts.Client().Get(MustParseURL(t, ts.URL, "healthz?ht=t"))
		if err != nil {
			t.Fatal(err)
		}

		if got, want := resp.StatusCode, http.StatusUnauthorized; got != want {
			t.Fatalf("want status code %d, got %d", want, got)
		}

		if got, want := MustReadBody(t, resp), "unauthorized\n"; got != want {
			t.Fatalf("want body %q, got %q", want, got)
		}
	})
}

func MustReadBody(t *testing.T, resp *http.Response) string {
	t.Helper()

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	return string(body)
}

func MustParseURL(t *testing.T, parts ...string) string {
	t.Helper()

	u, err := url.Parse(strings.Join(parts, "/"))
	if err != nil {
		t.Fatal(err)
	}
	return u.String()
}
