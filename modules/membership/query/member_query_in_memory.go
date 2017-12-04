package query

import (
	"errors"

	"github.com/purwokertodev/go-backend/modules/membership/model"
)

type memberQueryInMemory struct {
	db map[string]*model.Member
}

func NewMemberQueryInMemory(db map[string]*model.Member) MembershipQuery {
	return &memberQueryInMemory{db}
}

func (q *memberQueryInMemory) FindByEmail(email string) <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		var member *model.Member
		for _, v := range q.db {
			if v.Email == email {
				member = v
				break
			} else {
				output <- QueryResult{Error: errors.New("MEMBER_NOT_FOUND")}
				return
			}
		}

		output <- QueryResult{Result: member}
	}()
	return output
}
