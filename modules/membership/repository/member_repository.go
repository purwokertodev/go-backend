package repository

import (
	"github.com/purwokertodev/go-backend/modules/membership/model"
)

type RepositoryResult struct {
	Result interface{}
	Error  error
}

type MembershipRepository interface {
	Save(m *model.Member) <-chan error
	Load(id string) <-chan RepositoryResult
}
