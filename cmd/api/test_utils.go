package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/karthikbhandary2/Social/internal/auth"
	"github.com/karthikbhandary2/Social/internal/store"
	"github.com/karthikbhandary2/Social/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()

	// logger := zap.Must(zap.NewProduction()).Sugar()
	logger := zap.Must(zap.NewProduction()).Sugar()
	mockStore := store.NewMockStore()
	mockCacheStore := cache.NewMockStore()
	testAuth := &auth.TestAuthenticator{}
	return &application{
		logger: logger,
		store: mockStore,
		cacheStorage: mockCacheStore,
		authenticator: testAuth,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected status code %d, got %d", expected, actual)
	}
}