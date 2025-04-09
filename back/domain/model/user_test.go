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

func makeBirthday(val string) time.Time {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	birthday, _ := time.ParseInLocation("2006-01-02", val, jst)

	return birthday
}

func makeId(count int) string {
	id := ""
	for i := 0; i < count; i++ {
		id += "a"
	}

	return id
}

func makeName(count int) string {
	name := ""
	for i := 0; i < count; i++ {
		name += "a"
	}

	return name
}

func TestUser(t *testing.T) {

	cases := map[string]struct {
		id               string
		name             string
		birthday         time.Time
		email            *Email
		organizationName string
		isError          bool
	}{
		"success": {
			id:               makeId(50),
			name:             makeName(50),
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: makeId(50),
			isError:          false,
		},
		"emptyId": {
			id:               "",
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			isError:          true,
		},
		"overflowId": {
			id:               makeId(51),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			isError:          true,
		},
		"irregularCharacterId": {
			id:               "test@",
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			isError:          true,
		},
		"emptyName": {
			id:               makeId(50),
			name:             "",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			isError:          true,
		},
		"overflowName": {
			id:               makeId(50),
			name:             makeName(51),
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			isError:          true,
		},
		"irregularCharacterName": {
			id:               makeId(50),
			name:             "test@",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			isError:          true,
		},
		"emptyEmail": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            nil,
			organizationName: "testCompany",
			isError:          true,
		},
		"futureBirthday": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("3000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			isError:          true,
		},
		"overflowOrganizationName": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: makeId(51),
			isError:          true,
		},
		"emptyOrganizationName": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "",
			isError:          false,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			user, err := NewUser(
				c.id,
				c.name,
				c.birthday,
				c.email,
				"",
				"",
				"",
				c.organizationName,
				time.Now(),
				time.Now(),
				[]*Skill{},
				[]*ExternalServiceUrl{},
			)

			if c.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, c.id, user.ID())
			assert.Equal(t, c.name, user.Name())
			assert.Equal(t, c.birthday, user.Birthday())
			assert.Equal(t, c.email.Email(), user.Email().Email())
		})
	}
}
