package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMember(t *testing.T) {

	t.Run("New Member", func(t *testing.T) {
		m := NewMember()
		m.FirstName = "Wuriyanto"
		m.LastName = "Musobar"
		m.Email = "wuriyanto48@yahoo.co.id"
		m.Password = "12345"
		m.BirthDate = time.Now()

		assert.NotNil(t, m)
		assert.Equal(t, "Wuriyanto", m.FirstName)
		assert.Equal(t, "Musobar", m.LastName)
		assert.Equal(t, 0, m.Version)
	})

	t.Run("Test Valid Password", func(t *testing.T) {
		m := NewMember()
		m.FirstName = "Wuriyanto"
		m.LastName = "Musobar"
		m.Email = "wuriyanto48@yahoo.co.id"
		m.Password = "12345"
		m.BirthDate = time.Now()

		err := m.IsValidPassword("12345")
		assert.NoError(t, err)
	})

	t.Run("Test Invalid Password", func(t *testing.T) {
		m := NewMember()
		m.FirstName = "Wuriyanto"
		m.LastName = "Musobar"
		m.Email = "wuriyanto48@yahoo.co.id"
		m.Password = "12345"
		m.BirthDate = time.Now()

		err := m.IsValidPassword("123456")
		assert.Error(t, err)
	})
}
