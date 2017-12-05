package query

import (
	"errors"

	"github.com/purwokertodev/go-backend/modules/auth/model"
)

type identityQueryInMemory struct {
	db map[string]*model.Identity
}

func NewIdentityQueryInMemory(db map[string]*model.Identity) IdentityQuery {
	return &identityQueryInMemory{db}
}

func (q *identityQueryInMemory) FindByEmail(email string) <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		var identity *model.Identity
		for _, v := range q.db {
			if v.Email == email {
				identity = v
				break
			} else {
				output <- QueryResult{Error: errors.New("IDENTITY_NOT_FOUND")}
				return
			}
		}

		output <- QueryResult{Result: identity}
	}()
	return output
}
