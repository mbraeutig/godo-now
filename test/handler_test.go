package test

import (
	"net/http"
	"testing"
)

var _ http.ResponseWriter = (*TestResponseWriter)(nil)

type TestResponseWriter struct {
	header http.Header
}

func (t TestResponseWriter) Header() http.Header {
	return t.header
}

func (TestResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (TestResponseWriter) WriteHeader(statusCode int) {
}

func TestCreate(t *testing.T) {
}
