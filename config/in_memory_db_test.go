package config

import (
	"testing"

	"github.com/purwokertodev/go-backend/modules/membership/model"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryDB(t *testing.T) {
	db := GetInMemoryDb()

	wury := db["M1"]

	assert.Equal(t, "Wuriyanto", wury.FirstName)

	newData := model.NewMember()
	newData.ID = "M2"
	newData.FirstName = "Rob"
	newData.LastName = "Pike"

	db[newData.ID] = newData

	assert.Equal(t, "Rob", db["M2"].FirstName)
	assert.Equal(t, "Pike", db["M2"].LastName)

	assert.Equal(t, 2, len(db))
}
