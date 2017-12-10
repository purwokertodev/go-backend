package presenter

import (
	"net/http"

	"github.com/purwokertodev/go-backend/modules/membership/model"
	"github.com/purwokertodev/go-backend/modules/membership/usecase"
	"github.com/purwokertodev/go-backend/utils"
)

// HttpMembershipHandler model
type HttpMembershipHandler struct {
	memberUseCase usecase.MemberUseCase
}

// NewHttpMemberHandler for initialise HttpMembershipHandler model
func NewHttpMemberHandler(memberUseCase usecase.MemberUseCase) *HttpMembershipHandler {
	return &HttpMembershipHandler{memberUseCase: memberUseCase}
}

// Me http handler function, for get Member by its ID from Authorization
func (h *HttpMembershipHandler) Me() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {
			utils.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		memberID := req.Header.Get("MemberId")

		memberResult := <-h.memberUseCase.FindByID(memberID)

		if memberResult.Error != nil {
			utils.JsonResponse(res, "Member not found", http.StatusInternalServerError)
			return
		}

		member, ok := memberResult.Result.(*model.Member)

		if !ok {
			utils.JsonResponse(res, "Result is not member", http.StatusInternalServerError)
			return
		}

		utils.JsonResponse(res, member, http.StatusOK)

	})
}
