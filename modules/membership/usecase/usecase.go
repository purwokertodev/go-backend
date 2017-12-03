package usecase

import (
	"github.com/purwokertodev/go-backend/modules/membership/model"
)

type UseCaseResult struct {
	Result interface{}
	Error  error
}

type MemberUseCase interface {
	Save(m *model.Member) <-chan error
	FindByID(id string) <-chan UseCaseResult
	FindByEmail(email string) <-chan UseCaseResult
}
