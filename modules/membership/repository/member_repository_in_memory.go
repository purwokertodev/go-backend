package repository

import (
	"errors"
	"github.com/purwokertodev/go-backend/modules/membership/model"
	"time"
)

type MemberRepositoryInMemory struct {
	db map[string]*model.Member
}

func NewMemberRepositoryInMemory() *MemberRepositoryInMemory {
	return &MemberRepositoryInMemory{make(map[string]*model.Member)}
}

func (r *MemberRepositoryInMemory) Save(m *model.Member) <-chan error {
	output := make(chan error)
	go func() {
		member, ok := r.db[m.ID]
		if !ok {
			m.Version++
			r.db[m.ID] = m
			output <- nil
			return
		} else {
			member.Version++
			member.UpdatedAt = time.Now()
			r.db[m.ID] = member
			output <- nil
			return
		}
	}()
	return output
}

func (r *MemberRepositoryInMemory) Load(id string) <-chan RepositoryResult {
	output := make(chan RepositoryResult)
	go func() {
		member, ok := r.db[id]
		if !ok {
			output <- RepositoryResult{Error: errors.New("member not found")}
			return
		}

		output <- RepositoryResult{Result: member}
	}()
	return output
}
