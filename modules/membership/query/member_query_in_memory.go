package query

import (
	"errors"

	"github.com/purwokertodev/go-backend/modules/membership/model"
)

type MemberQueryInMemory struct {
	db map[string]*model.Member
}

func NewMemberQueryInMemory(db map[string]*model.Member) *MemberQueryInMemory {
	return &MemberQueryInMemory{db}
}

func (q *MemberQueryInMemory) FindByEmail(email string) <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		var member *model.Member
		for _, v := range q.db {
			if v.Email == email {
				member = v
				break
			} else {
				output <- QueryResult{Error: errors.New("Member not found")}
				return
			}
		}

		output <- QueryResult{Result: member}
	}()
	return output
}
