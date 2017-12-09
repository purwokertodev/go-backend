package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/purwokertodev/go-backend/modules/auth/model"
	"github.com/purwokertodev/go-backend/modules/auth/token"
	"github.com/purwokertodev/go-backend/modules/auth/usecase"
)

type HttpAuthHandler struct {
	AuthUseCase usecase.AuthUseCase
}

func NewHttpHandler(authUseCase usecase.AuthUseCase) *HttpAuthHandler {
	return &HttpAuthHandler{AuthUseCase: authUseCase}
}

func (h *HttpAuthHandler) Auth() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "POST" {
			jsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		grantTypes, ok := req.URL.Query()["grant_type"]

		if !ok || len(grantTypes) < 1 {
			jsonResponse(res, "Missing Grant Type", http.StatusBadRequest)
			return
		}

		// Query()["grant_type"] will return an array of items,
		// we only want the single item.
		grantType := grantTypes[0]

		switch grantType {
		case "password":
			var identityLogin model.Identity

			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&identityLogin)

			if err != nil {
				jsonResponse(res, "Error occured", http.StatusInternalServerError)
				return
			}

			accessTokenResult := <-h.AuthUseCase.GetAccessToken(identityLogin.Email, identityLogin.Password)

			if accessTokenResult.Error != nil {
				jsonResponse(res, "Invalid username or password", http.StatusUnauthorized)
				return
			}

			accessToken, ok := accessTokenResult.Result.(*token.AccessToken)

			if !ok {
				jsonResponse(res, "Invalid username or password", http.StatusUnauthorized)
				return
			}

			jsonResponse(res, accessToken, http.StatusOK)
			return
		default:
			jsonResponse(res, "Invalid Grant Type", http.StatusBadRequest)
			return
		}
	})
}

func jsonResponse(res http.ResponseWriter, resp interface{}, httpCode int) {
	msg, _ := json.Marshal(resp)
	res.Header().Set("Content-Type", "application-json; charset=utf-8")
	res.WriteHeader(httpCode)
	res.Write(msg)
}
