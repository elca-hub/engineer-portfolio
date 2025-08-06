package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeSkillName(count int) string {
	name := ""
	for i := 0; i < count; i++ {
		name += "a"
	}
	return name
}

func makeSkillComment(count int) string {
	comment := ""
	for i := 0; i < count; i++ {
		comment += "a"
	}
	return comment
}

func makeSkillYear(val string) time.Time {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	year, _ := time.ParseInLocation("2006-01-02", val, jst)
	return year
}

func TestNewSkill(t *testing.T) {
	cases := map[string]struct {
		name      string
		year      time.Time
		comment   string
		sortIndex int
		isError   bool
	}{
		"success": {
			name:      makeSkillName(MaxSkillNameLength),
			year:      makeSkillYear("2020-01-01"),
			comment:   makeSkillComment(MaxSkillCommentLength),
			sortIndex: 1,
			isError:   false,
		},
		"emptyName": {
			name:      "",
			year:      makeSkillYear("2020-01-01"),
			comment:   "test comment",
			sortIndex: 1,
			isError:   true,
		},
		"overflowName": {
			name:      makeSkillName(MaxSkillNameLength + 1),
			year:      makeSkillYear("2020-01-01"),
			comment:   "test comment",
			sortIndex: 1,
			isError:   true,
		},
		"emptyYear": {
			name:      "test skill",
			year:      time.Time{},
			comment:   "test comment",
			sortIndex: 1,
			isError:   true,
		},
		"overflowComment": {
			name:      "test skill",
			year:      makeSkillYear("2020-01-01"),
			comment:   makeSkillComment(MaxSkillCommentLength + 1),
			sortIndex: 1,
			isError:   true,
		},
		"emptyComment": {
			name:      "test skill",
			year:      makeSkillYear("2020-01-01"),
			comment:   "",
			sortIndex: 1,
			isError:   false,
		},
		"negativeSortIndex": {
			name:      "test skill",
			year:      makeSkillYear("2020-01-01"),
			comment:   "test comment",
			sortIndex: -1,
			isError:   true,
		},
		"zeroSortIndex": {
			name:      "test skill",
			year:      makeSkillYear("2020-01-01"),
			comment:   "test comment",
			sortIndex: 0,
			isError:   false,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			skill, err := NewSkill(c.name, c.year, c.comment, c.sortIndex)

			if c.isError {
				assert.Error(t, err)
				assert.Nil(t, skill)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, skill)
			assert.Equal(t, c.name, skill.Name())
			assert.Equal(t, c.year.Truncate(24*time.Hour), skill.Year())
			assert.Equal(t, c.comment, skill.Comment())
			assert.Equal(t, c.sortIndex, skill.SortIndex())
		})
	}
}

func TestSkill_UpdateName(t *testing.T) {
	skill, _ := NewSkill("original name", makeSkillYear("2020-01-01"), "comment", 1)

	cases := map[string]struct {
		newName string
		isError bool
	}{
		"success": {
			newName: "updated name",
			isError: false,
		},
		"emptyName": {
			newName: "",
			isError: true,
		},
		"overflowName": {
			newName: makeSkillName(MaxSkillNameLength + 1),
			isError: true,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			testSkill := &Skill{
				name:      skill.name,
				year:      skill.year,
				comment:   skill.comment,
				sortIndex: skill.sortIndex,
			}

			err := testSkill.UpdateName(c.newName)

			if c.isError {
				assert.Error(t, err)
				assert.Equal(t, skill.name, testSkill.name) // 元の値が変わっていないことを確認
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, c.newName, testSkill.Name())
		})
	}
}

func TestSkill_UpdateYear(t *testing.T) {
	skill, _ := NewSkill("skill name", makeSkillYear("2020-01-01"), "comment", 1)

	cases := map[string]struct {
		newYear string
		isError bool
	}{
		"success": {
			newYear: "2021-01-01",
			isError: false,
		},
		"emptyYear": {
			newYear: "",
			isError: true,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			testSkill := &Skill{
				name:      skill.name,
				year:      skill.year,
				comment:   skill.comment,
				sortIndex: skill.sortIndex,
			}

			var newYearTime time.Time
			if c.newYear != "" {
				newYearTime = makeSkillYear(c.newYear)
			}

			err := testSkill.UpdateYear(newYearTime)

			if c.isError {
				assert.Error(t, err)
				assert.Equal(t, skill.year, testSkill.year) // 元の値が変わっていないことを確認
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, newYearTime.Truncate(24*time.Hour), testSkill.Year())
		})
	}
}

func TestSkill_UpdateComment(t *testing.T) {
	skill, _ := NewSkill("skill name", makeSkillYear("2020-01-01"), "original comment", 1)

	cases := map[string]struct {
		newComment string
		isError    bool
	}{
		"success": {
			newComment: "updated comment",
			isError:    false,
		},
		"emptyComment": {
			newComment: "",
			isError:    false,
		},
		"overflowComment": {
			newComment: makeSkillComment(MaxSkillCommentLength + 1),
			isError:    true,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			testSkill := &Skill{
				name:      skill.name,
				year:      skill.year,
				comment:   skill.comment,
				sortIndex: skill.sortIndex,
			}

			err := testSkill.UpdateComment(c.newComment)

			if c.isError {
				assert.Error(t, err)
				assert.Equal(t, skill.comment, testSkill.comment) // 元の値が変わっていないことを確認
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, c.newComment, testSkill.Comment())
		})
	}
}

func TestSkill_UpdateSortIndex(t *testing.T) {
	skill, _ := NewSkill("skill name", makeSkillYear("2020-01-01"), "comment", 1)

	cases := map[string]struct {
		newSortIndex int
		isError      bool
	}{
		"success": {
			newSortIndex: 5,
			isError:      false,
		},
		"zeroSortIndex": {
			newSortIndex: 0,
			isError:      false,
		},
		"negativeSortIndex": {
			newSortIndex: -1,
			isError:      true,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			testSkill := &Skill{
				name:      skill.name,
				year:      skill.year,
				comment:   skill.comment,
				sortIndex: skill.sortIndex,
			}

			err := testSkill.UpdateSortIndex(c.newSortIndex)

			if c.isError {
				assert.Error(t, err)
				assert.Equal(t, skill.sortIndex, testSkill.sortIndex) // 元の値が変わっていないことを確認
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, c.newSortIndex, testSkill.SortIndex())
		})
	}
}