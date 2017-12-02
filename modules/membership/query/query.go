package query

import (
	"github.com/purwokertodev/go-backend/modules/membership/model"
)

type QueryResult struct {
	Result interface{}
	Error  error
}

type MemberQuery interface {
	FindByEmail(email string) <-chan QueryResult
}
