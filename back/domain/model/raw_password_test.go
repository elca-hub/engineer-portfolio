package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_, err := NewRawPassword("Security_1234")

		assert.NoError(t, err)
	})

	t.Run("failures", func(t *testing.T) {
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
				_, err := NewRawPassword(c.password)

				assert.Error(t, err)
			})
		}
	})
}
