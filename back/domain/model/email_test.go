package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrueEmail(t *testing.T) {
	_, err := NewEmail("test@example.com")

	assert.NoError(t, err)
}

func TestFalseEmail(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		_, err := NewEmail("test@example")

		assert.Error(t, err)
	})

	t.Run("Empty", func(t *testing.T) {
		_, err := NewEmail("")

		assert.Error(t, err)
	})
}
