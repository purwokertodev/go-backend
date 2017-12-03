package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {

	req := httptest.NewRequest("GET", "/cek", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "application/json")

		io.WriteString(res, `{"alive": true}`)
	})

	handler := LogRequest(testHandler)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

}
