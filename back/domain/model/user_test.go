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

func TestUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_, err := NewUser(
			NewUUID(""),
			"test",
			0,
			fetchEmail(),
			"Security_3939",
			time.Now(),
			time.Now(),
			Unconfirmed,
		)

		assert.NoError(t, err)
	})

	t.Run("failures", func(t *testing.T) {
		t.Run("Name", func(t *testing.T) {
			tooLongName := ""

			for i := 0; i < MaxNameLen+1; i++ {
				tooLongName += "a"
			}

			cases := map[string]struct {
				name string
			}{
				"empty": {
					name: "",
				},
				"too long": {
					name: tooLongName,
				},
				"special characters": {
					name: "test@",
				},
				"space": {
					name: "test test",
				},
			}

			for name, c := range cases {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					_, err := NewUser(
						NewUUID(""),
						c.name,
						0,
						fetchEmail(),
						"Security_3939",
						time.Now(),
						time.Now(),
						Unconfirmed,
					)

					assert.Error(t, err)
				})
			}
		})

		t.Run("Age", func(t *testing.T) {
			cases := map[string]struct {
				age int
			}{
				"negative": {
					age: -1,
				},
			}

			for name, c := range cases {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					_, err := NewUser(
						NewUUID(""),
						"test",
						c.age,
						fetchEmail(),
						"Security_3939",
						time.Now(),
						time.Now(),
						Unconfirmed,
					)

					assert.Error(t, err)
				})
			}
		})

		t.Run("Password", func(t *testing.T) {
			tooShortPassword := ""
			for i := 0; i < MinPasswordLen-1; i++ {
				tooShortPassword += "a"
			}

			tooLongPassword := ""
			for i := 0; i < MaxPasswordLen+1; i++ {
				tooLongPassword += "a"
			}

			cases := map[string]struct {
				password string
			}{
				"empty": {
					password: "",
				},
				"too short": {
					password: tooShortPassword,
				},
				"too long": {
					password: tooLongPassword,
				},
				"not using special characters": {
					password: "security3939",
				},
				"not using numbers": {
					password: "security_",
				},
				"not using lowercase letters": {
					password: "SECURITY_3939",
				},
				"not using uppercase letters": {
					password: "security_3939",
				},
			}

			for name, c := range cases {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					_, err := NewUser(
						NewUUID(""),
						"test",
						0,
						fetchEmail(),
						c.password,
						time.Now(),
						time.Now(),
						Unconfirmed,
					)

					assert.Error(t, err)
				})
			}
		})
	})
}
