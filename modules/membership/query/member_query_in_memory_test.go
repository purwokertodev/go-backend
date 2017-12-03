package query

import (
	"testing"
	"time"

	"github.com/purwokertodev/go-backend/modules/membership/model"
	"github.com/stretchr/testify/assert"
)

func TestMemberRepositoryInMemory(t *testing.T) {

	db := make(map[string]*model.Member)

	exampeMember := model.NewMember()
	exampeMember.ID = "M1"
	exampeMember.FirstName = "Wuriyanto"
	exampeMember.LastName = "Musobar"
	exampeMember.Email = "wuriyanto48@yahoo.co.id"
	exampeMember.Password = "123456"
	exampeMember.PasswordSalt = "salt"
	exampeMember.BirthDate = time.Now()

	db["M1"] = exampeMember

	q := NewMemberQueryInMemory(db)

	t.Run("Find Member By Email", func(t *testing.T) {

		memberResult := <-q.FindByEmail("wuriyanto48@yahoo.co.id")

		assert.NoError(t, memberResult.Error)

		wury, ok := memberResult.Result.(*model.Member)

		assert.True(t, ok)
		assert.Equal(t, "Wuriyanto", wury.FirstName)
		assert.Equal(t, "Musobar", wury.LastName)
		assert.Equal(t, 0, wury.Version)

	})

	t.Run("Find Member By Email Not Found", func(t *testing.T) {

		memberResult := <-q.FindByEmail("invalid@email.com")

		assert.Error(t, memberResult.Error)

		_, ok := memberResult.Result.(*model.Member)
		assert.False(t, ok)

	})

}
