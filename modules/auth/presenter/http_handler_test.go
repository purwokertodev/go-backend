package presenter

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/purwokertodev/go-backend/modules/auth/model"
	"github.com/purwokertodev/go-backend/modules/auth/query"
	"github.com/purwokertodev/go-backend/modules/auth/token"
	"github.com/purwokertodev/go-backend/modules/auth/usecase"
	"github.com/purwokertodev/go-backend/modules/auth/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func generateIdentityAccessTokenResult() <-chan usecase.UseCaseResult {
	output := make(chan usecase.UseCaseResult)
	go func() {
		var accessToken token.AccessToken
		accessToken.Token = `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJpYXQiOjE1MTIzMTg0MjMsImlzcyI6InB1cndva2VydG9kZXYuZ2l0aHViLmlvIiwic3ViIjoiMDAxIn0.GeEDHv82F8xp_98QQLiWxZ5aVBzZej0e-Ios8M9l0tYdrrTbdP3zxutSi5H7rxrd43PmlFi0pMMGtbVw64kPkBspCE3Kebbeersa8isn1zBejZO62mIgpRIGRhAJ_rphsxXYqOKKQlgj2ecI39dRR7IRJZdNYoTXXeBktUeUcDU`

		output <- usecase.UseCaseResult{Result: &accessToken}
	}()
	return output
}

func generateQueryFindByEmailSuccessIdentityResult() <-chan query.QueryResult {
	output := make(chan query.QueryResult)
	go func() {
		var i model.Identity
		i.ID = "M1"
		i.Email = "wuriyanto48@yahoo.co.id"
		i.Password = "12345"
		output <- query.QueryResult{Result: i}
	}()
	return output
}

func TestAuthHandler(t *testing.T) {

	var i model.Identity
	i.ID = "M1"
	i.Email = "wuriyanto48@yahoo.co.id"
	i.Password = "12345"

	t.Run("Test Password Credentials", func(t *testing.T) {

		mockUseCase := new(mocks.AuthUseCase)

		body, _ := json.Marshal(i)

		mockUseCase.On("GetAccessToken", mock.AnythingOfType("string,string")).Return(generateIdentityAccessTokenResult())
		req := httptest.NewRequest("POST", "/auth?grant_type=password", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		//signKey, _ := getPrivateKey(PrivateKey)

		httpAuthHandler := NewHttpHandler(mockUseCase)

		handler := http.Handler(httpAuthHandler.Auth())

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.

		handler.ServeHTTP(rec, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusOK, rec.Code)

	})
}

const PrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCoqzL5JrMzed4tb8uEoLKd42EOsYmb0HpbicGt/OUeJxaHtt59
Ew0BbpreBeiuugXweEa5xctQOxGYr27h4ZOnR0hWSi+h5Y35CKzMEmZnzQwzQphg
qww0U+e9/OAvVfCW1xWvVFr0WbhIRn+w/9DUvp+6jKz3fIj3yQaHWVMMNQIDAQAB
AoGAKd7d94XI5JVzNxpSjmkKDjHc7TXbcEevqDupTdTC19piOGyIDMqG5v0bCtSy
r3VUdh6ViBZ240LWmm2qe/5wlaUSDxGAhQg78cVP9L157hC15vOOckMjcJyuVCpR
Wew61HP3JNPB3dsk8P/fjPwgXzsH9L0zIoT0Krw85TbY8UECQQDsRoyWMPu5V6Sa
kiQrr2hZ+weCRH9Q6yd97UhxPAgSZZswedn97uF5T3GdKvoLpeyUrpBe7x3Y3Ciz
uN+SGBPlAkEAtr/SncRbSqUbdrc5BcAVnFNwLXJc2bN477z7shs0OIAha0rcq/8Q
jC5M5oh8jIM5bltZ5t8CWHbrfryVSwsyEQJACnkiGDI5pkCNSlC6C7mtvXdUIOEa
Z6LU0E8pS+OmU/JvC5oLIKdrFS6BUb8q8EM9lmWafqrIvukbYMQMHPS2RQJASWYN
35PH3tkliK7aVjbp9xmECpzOMhnlTtSmesh2VuMPiRpOOz58lPDbrhPPglgKLwq9
tv6G4KUSvJpdlABxIQJARw4I/XUNdd2ko+gSkEHjwjKg4LlYNydHHGk6RYc1S85L
U8PhZfO17Ul2d9ROvFHx75slASSgWHEnPUF7gphhUA==
-----END RSA PRIVATE KEY-----
`

func getPrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	r := strings.NewReader(privateKey)
	signBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}
	return signKey, nil
}
