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
		occupationName   string
		place            string
		isError          bool
	}{
		"success": {
			id:               makeId(50),
			name:             makeName(50),
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: makeId(50),
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          false,
		},
		"emptyId": {
			id:               "",
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          true,
		},
		"overflowId": {
			id:               makeId(51),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          true,
		},
		"irregularCharacterId": {
			id:               "test@",
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          true,
		},
		"emptyName": {
			id:               makeId(50),
			name:             "",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          true,
		},
		"overflowName": {
			id:               makeId(50),
			name:             makeName(51),
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          true,
		},
		"irregularCharacterName": {
			id:               makeId(50),
			name:             "test@",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          true,
		},
		"emptyEmail": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            nil,
			organizationName: "testCompany",
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          true,
		},
		"futureBirthday": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("3000-01-01"),
			email:            fetchEmail(),
			organizationName: "testCompany",
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          true,
		},
		"overflowOrganizationName": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: makeId(51),
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          true,
		},
		"emptyOrganizationName": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: "",
			occupationName:   makeId(50),
			place:            makeId(50),
			isError:          false,
		},
		"overflowOccupationName": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: makeId(50),
			occupationName:   makeId(51),
			isError:          true,
		},
		"emptyOccupationName": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: makeId(50),
			occupationName:   "",
			place:            makeId(50),
			isError:          false,
		},
		"overflowPlace": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: makeId(50),
			occupationName:   makeId(50),
			place:            makeId(51),
			isError:          true,
		},
		"emptyPlace": {
			id:               makeId(50),
			name:             "test",
			birthday:         makeBirthday("2000-01-01"),
			email:            fetchEmail(),
			organizationName: makeId(50),
			occupationName:   makeId(50),
			place:            "",
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
				nil,
				c.organizationName,
				c.occupationName,
				c.place,
				time.Now(),
				time.Now(),
				[]*Skill{},
				[]*Certification{},
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

func TestUser_UpdateSkills(t *testing.T) {
	user, _ := NewUser(
		"testuser",
		"test",
		makeBirthday("2000-01-01"),
		fetchEmail(),
		"",
		"",
		nil,
		"testCompany",
		"engineer",
		"tokyo",
		time.Now(),
		time.Now(),
		[]*Skill{},
		[]*Certification{},
		[]*ExternalServiceUrl{},
	)

	skill1, _ := NewSkill("Go", makeSkillYear("2020-01-01"), "Programming language", 1)
	skill2, _ := NewSkill("Python", makeSkillYear("2019-01-01"), "Programming language", 2)

	cases := map[string]struct {
		skills  []*Skill
		isError bool
	}{
		"success": {
			skills:  []*Skill{skill1, skill2},
			isError: false,
		},
		"emptySkills": {
			skills:  []*Skill{},
			isError: false,
		},
		"tooManySkills": {
			skills:  make([]*Skill, MaxUserSkillsLen+1),
			isError: true,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := user.UpdateSkills(c.skills)

			if c.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, len(c.skills), len(user.Skills()))
			for i, skill := range c.skills {
				assert.Equal(t, skill.Name(), user.Skills()[i].Name())
			}
		})
	}
}

func TestUser_UpdateCertifications(t *testing.T) {
	user, _ := NewUser(
		"testuser",
		"test",
		makeBirthday("2000-01-01"),
		fetchEmail(),
		"",
		"",
		nil,
		"testCompany",
		"engineer",
		"tokyo",
		time.Now(),
		time.Now(),
		[]*Skill{},
		[]*Certification{},
		[]*ExternalServiceUrl{},
	)

	cert1, _ := NewCertification("AWS Solutions Architect", makeCertYear("2021-01-01"), "Cloud certification", 1)
	cert2, _ := NewCertification("TOEIC 900", makeCertYear("2020-01-01"), "English certification", 2)

	cases := map[string]struct {
		certifications []*Certification
		isError        bool
	}{
		"success": {
			certifications: []*Certification{cert1, cert2},
			isError:        false,
		},
		"emptyCertifications": {
			certifications: []*Certification{},
			isError:        false,
		},
		"tooManyCertifications": {
			certifications: make([]*Certification, MaxUserSkillsLen+1),
			isError:        true,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := user.UpdateCertifications(c.certifications)

			if c.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, len(c.certifications), len(user.Certifications()))
			for i, cert := range c.certifications {
				assert.Equal(t, cert.Name(), user.Certifications()[i].Name())
			}
		})
	}
}


