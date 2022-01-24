// Copyright 2020 Eurac Research. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package http handles everything related to HTTP.
package http

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"

	"github.com/euracresearch/browser"
	"golang.org/x/crypto/acme/autocert"
)

const languageCookieName = "browser_lter_lang"

// ListenAndServe is a wrapper for http.ListenAndServe.
func ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}

// ServeAutoCert will serve on the standard TLS port (443) with
// LetsEncrypt certificates for the provided domain or domains.
// Certificates will be stored in the given cache directory. Incoming
// traffic on port 80 will be automatically forwared to 443.
func ServeAutoCert(addr string, handler http.Handler, cache string, domains ...string) error {
	go func() {
		host, _, err := net.SplitHostPort(addr)
		if err != nil || host == "" {
			host = "0.0.0.0"
		}
		log.Println("Redirecting traffic from HTTP to HTTPS.")
		log.Fatal(http.ListenAndServe(host+":80", redirectHandler()))
	}()

	m := &autocert.Manager{
		Cache:      autocert.DirCache(cache),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domains...),
	}

	m.TLSConfig().MinVersion = tls.VersionTLS12
	m.TLSConfig().CurvePreferences = []tls.CurveID{
		tls.CurveP256,
		tls.X25519, // Go 1.8 only
	}
	m.TLSConfig().CipherSuites = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	}

	s := &http.Server{
		Addr:      addr,
		Handler:   handler,
		TLSConfig: m.TLSConfig(),
	}

	return s.ListenAndServeTLS("", "")
}

func redirectHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "close")
		url := "https://" + r.Host + r.URL.String()
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	})
}

// Error writes an error message to the response.
func Error(w http.ResponseWriter, err error, code int) {
	// Log error.
	log.Printf("http error: %s (code=%d)", err, code)

	// Hide error message from client if it is internal or not found.
	if code == http.StatusInternalServerError || code == http.StatusNotFound {
		err = browser.ErrInternal
	}

	http.Error(w, err.Error(), code)
}

// grantAccess is a HTTP middleware function which grants access to the given
// handler to the given roles.
func grantAccess(h http.HandlerFunc, roles ...browser.Role) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAllowed(r, roles...) {
			http.NotFound(w, r)
			return
		}

		h(w, r)
	}
}

// isAllowed checks if the current user makes part of the allowed roles.
func isAllowed(r *http.Request, roles ...browser.Role) bool {
	u := browser.UserFromContext(r.Context())

	for _, v := range roles {
		if u.Role == v {
			return true
		}
	}

	return false
}
