package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMember(t *testing.T) {

	t.Run("New Member", func(t *testing.T) {
		m := NewMember()
		m.FirstName = "Wuriyanto"
		m.LastName = "Musobar"
		m.BirthDate = time.Now()

		assert.NotNil(t, m)
		assert.Equal(t, "Wuriyanto", m.FirstName)
		assert.Equal(t, "Musobar", m.LastName)
		assert.Equal(t, 0, m.Version)
	})
}
