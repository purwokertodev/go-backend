package config

import (
	"time"

	"github.com/purwokertodev/go-backend/modules/membership/model"
)

func GetInMemroryDb() map[string]*model.Member {
	db := make(map[string]*model.Member)

	exampeMember := model.NewMember()
	exampeMember.ID = "M1"
	exampeMember.FirstName = "Wuriyanto"
	exampeMember.LastName = "Musobar"
	exampeMember.Email = "wuriyanto48@yahoo.co.id"
	exampeMember.Password = "123456"
	exampeMember.PasswordSalt = "salt"
	exampeMember.BirthDate = time.Now()

	db[exampeMember.ID] = exampeMember

	return db
}
