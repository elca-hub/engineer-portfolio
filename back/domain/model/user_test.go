package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func fetchEmail() *Email {
	email, _ := NewEmail("test@example.com")

	return email
}

const (
	name     = "test"
	age      = 18
	password = "password"
)

func TestTrueUser(t *testing.T) {
	_, err := NewUser(
		NewUUID(""),
		"test",
		0,
		fetchEmail(),
		"password",
		time.Now(),
		time.Now(),
		Unconfirmed,
	)

	assert.NoError(t, err)
}

func TestFalseUser(t *testing.T) {
	t.Run("Name", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			_, err := NewUser(
				NewUUID(""),
				"",
				0,
				fetchEmail(),
				"password",
				time.Now(),
				time.Now(),
				Unconfirmed,
			)

			assert.Error(t, err)
		})

		t.Run("too long", func(t *testing.T) {
			longName := ""

			for i := 0; i < MaxNameLen; i++ {
				longName += "a"
			}
			_, err := NewUser(
				NewUUID(""),
				longName,
				0,
				fetchEmail(),
				"password",
				time.Now(),
				time.Now(),
				Unconfirmed,
			)

			assert.Error(t, err)
		})

		t.Run("use space", func(t *testing.T) {
			_, err := NewUser(
				NewUUID(""),
				"test test",
				0,
				fetchEmail(),
				"password",
				time.Now(),
				time.Now(),
				Unconfirmed,
			)

			assert.Error(t, err)
		})

		t.Run("use special characters", func(t *testing.T) {
			_, err := NewUser(
				NewUUID(""),
				"\"test\n@.,'",
				0,
				fetchEmail(),
				"password",
				time.Now(),
				time.Now(),
				Unconfirmed,
			)

			assert.Error(t, err)
		})
	})
}
