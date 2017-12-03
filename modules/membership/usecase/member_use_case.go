package usecase

import (
	"errors"

	"github.com/purwokertodev/go-backend/modules/membership/model"
	"github.com/purwokertodev/go-backend/modules/membership/query"
	"github.com/purwokertodev/go-backend/modules/membership/repository"
)

type memberUseCaseImpl struct {
	memberRepository repository.MembershipRepository
	memberQuery      query.MembershipQuery
}

func NewMemberUseCase(memberRepository repository.MembershipRepository, memberQuery query.MembershipQuery) MemberUseCase {
	return &memberUseCaseImpl{
		memberRepository: memberRepository,
		memberQuery:      memberQuery,
	}
}

func (mu *memberUseCaseImpl) Save(m *model.Member) <-chan error {
	output := make(chan error)

	go func() {

		memberExistResult := <-mu.memberQuery.FindByEmail(m.Email)

		if memberExistResult.Error != nil {
			output <- memberExistResult.Error
			return
		}

		memberExist, ok := memberExistResult.Result.(*model.Member)

		if !ok {
			output <- errors.New("Result is not member")
			return
		}

		if memberExist != nil {
			output <- errors.New("Member aleady exist")
			return
		}

		err := <-mu.memberRepository.Save(m)

		if err != nil {
			output <- err
			return
		}

	}()

	return output
}

func (mu *memberUseCaseImpl) FindByID(id string) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		memberResult := <-mu.memberRepository.Load(id)

		if memberResult.Error != nil {
			output <- UseCaseResult{Error: memberResult.Error}
			return
		}

		member, ok := memberResult.Result.(*model.Member)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Result is not member")}
			return
		}

		output <- UseCaseResult{Result: member}

	}()

	return output
}

func (mu *memberUseCaseImpl) FindByEmail(email string) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		memberResult := <-mu.memberQuery.FindByEmail(email)

		if memberResult.Error != nil {
			output <- UseCaseResult{Error: memberResult.Error}
			return
		}

		member, ok := memberResult.Result.(*model.Member)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Result is not member")}
			return
		}

		output <- UseCaseResult{Result: member}

	}()

	return output
}
