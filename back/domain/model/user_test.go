package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func fetchEmail() *Email {
	email, _ := NewEmail("test@example.com")

	return email
}

func TestUser(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	birthday, _ := time.ParseInLocation("2006-01-02", "1990-01-01", jst)
	t.Run("success", func(t *testing.T) {

		_, err := NewUser(
			"test",
			"test",
			birthday,
			fetchEmail(),
			"",
			"",
			"",
			[],
			[],
		)

		assert.NoError(t, err)
	})

	t.Run("failures", func(t *testing.T) {
		t.Run("Name", func(t *testing.T) {
			tooLongName := ""

			for i := 0; i < MaxUserNameLen+1; i++ {
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
						"test",
						c.name,
						birthday,
						fetchEmail(),
						time.Now(),
						time.Now(),
					)

					assert.Error(t, err)
				})
			}
		})

		t.Run("Birthday", func(t *testing.T) {
			cases := map[string]struct {
				birthdayCase time.Time
			}{
				"negative": {
					birthdayCase: time.Now().AddDate(0, 0, 1),
				},
			}

			for name, c := range cases {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					_, err := NewUser(
						"test",
						"test",
						c.birthdayCase,
						fetchEmail(),
						time.Now(),
						time.Now(),
					)

					assert.Error(t, err)
				})
			}
		})
	})
}
