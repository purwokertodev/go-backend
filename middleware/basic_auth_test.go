package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuth(t *testing.T) {

	validBasicAuth := `Basic cHVyd29rZXJ0b2RldjoxMjM0NTY=`
	invalidBasicAuth := `Basic cHVyd29rZXJ0b2RldjoxMjM0NTY`
	invalidBasicAuthMissingHeader := `cHVyd29rZXJ0b2RldjoxMjM0NTY`

	t.Run("Test Valid Basic Auth", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", validBasicAuth)
		rec := httptest.NewRecorder()

		testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/json")

			io.WriteString(res, `{"alive": true}`)
		})

		conf := NewConfig("purwokertodev", "123456")

		handler := BasicAuth(conf, testHandler)
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Test Invalid Basic Auth", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", invalidBasicAuth)
		rec := httptest.NewRecorder()

		testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/json")

			io.WriteString(res, `{"alive": true}`)
		})

		conf := NewConfig("purwokertodev", "123456")

		handler := BasicAuth(conf, testHandler)
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("Test Invalid Basic Auth Missing Header", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", invalidBasicAuthMissingHeader)
		rec := httptest.NewRecorder()

		testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/json")

			io.WriteString(res, `{"alive": true}`)
		})

		conf := NewConfig("purwokertodev", "123456")

		handler := BasicAuth(conf, testHandler)
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})
}
