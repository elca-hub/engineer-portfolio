package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeCertificationName(count int) string {
	name := ""
	for i := 0; i < count; i++ {
		name += "a"
	}
	return name
}

func makeCertificationComment(count int) string {
	comment := ""
	for i := 0; i < count; i++ {
		comment += "a"
	}
	return comment
}

func makeCertYear(val string) time.Time {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	year, _ := time.ParseInLocation("2006-01-02", val, jst)
	return year
}

func TestNewCertification(t *testing.T) {
	cases := map[string]struct {
		name      string
		year      time.Time
		comment   string
		sortIndex int
		isError   bool
	}{
		"success": {
			name:      makeCertificationName(MaxCertificationNameLength),
			year:      makeCertYear("2020-01-01"),
			comment:   makeCertificationComment(MaxCertificationCommentLength),
			sortIndex: 1,
			isError:   false,
		},
		"emptyName": {
			name:      "",
			year:      makeCertYear("2020-01-01"),
			comment:   "test comment",
			sortIndex: 1,
			isError:   true,
		},
		"overflowName": {
			name:      makeCertificationName(MaxCertificationNameLength + 1),
			year:      makeCertYear("2020-01-01"),
			comment:   "test comment",
			sortIndex: 1,
			isError:   true,
		},
		"emptyYear": {
			name:      "test certification",
			year:      time.Time{},
			comment:   "test comment",
			sortIndex: 1,
			isError:   true,
		},
		"overflowComment": {
			name:      "test certification",
			year:      makeCertYear("2020-01-01"),
			comment:   makeCertificationComment(MaxCertificationCommentLength + 1),
			sortIndex: 1,
			isError:   true,
		},
		"emptyComment": {
			name:      "test certification",
			year:      makeCertYear("2020-01-01"),
			comment:   "",
			sortIndex: 1,
			isError:   false,
		},
		"negativeSortIndex": {
			name:      "test certification",
			year:      makeCertYear("2020-01-01"),
			comment:   "test comment",
			sortIndex: -1,
			isError:   true,
		},
		"zeroSortIndex": {
			name:      "test certification",
			year:      makeCertYear("2020-01-01"),
			comment:   "test comment",
			sortIndex: 0,
			isError:   false,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			certification, err := NewCertification(c.name, c.year, c.comment, c.sortIndex)

			if c.isError {
				assert.Error(t, err)
				assert.Nil(t, certification)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, certification)
			assert.Equal(t, c.name, certification.Name())
			assert.Equal(t, c.year.Truncate(24*time.Hour), certification.Year())
			assert.Equal(t, c.comment, certification.Comment())
			assert.Equal(t, c.sortIndex, certification.SortIndex())
		})
	}
}

func TestCertification_UpdateName(t *testing.T) {
	certification, _ := NewCertification("original name", makeCertYear("2020-01-01"), "comment", 1)

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
			newName: makeCertificationName(MaxCertificationNameLength + 1),
			isError: true,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			testCertification := &Certification{
				name:      certification.name,
				year:      certification.year,
				comment:   certification.comment,
				sortIndex: certification.sortIndex,
			}

			err := testCertification.UpdateName(c.newName)

			if c.isError {
				assert.Error(t, err)
				assert.Equal(t, certification.name, testCertification.name) // 元の値が変わっていないことを確認
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, c.newName, testCertification.Name())
		})
	}
}

func TestCertification_UpdateYear(t *testing.T) {
	certification, _ := NewCertification("certification name", makeCertYear("2020-01-01"), "comment", 1)

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
			testCertification := &Certification{
				name:      certification.name,
				year:      certification.year,
				comment:   certification.comment,
				sortIndex: certification.sortIndex,
			}

			var newYearTime time.Time
			if c.newYear != "" {
				newYearTime = makeCertYear(c.newYear)
			}

			err := testCertification.UpdateYear(newYearTime)

			if c.isError {
				assert.Error(t, err)
				assert.Equal(t, certification.year, testCertification.year) // 元の値が変わっていないことを確認
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, newYearTime.Truncate(24*time.Hour), testCertification.Year())
		})
	}
}

func TestCertification_UpdateComment(t *testing.T) {
	certification, _ := NewCertification("certification name", makeCertYear("2020-01-01"), "original comment", 1)

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
			newComment: makeCertificationComment(MaxCertificationCommentLength + 1),
			isError:    true,
		},
	}

	t.Parallel()
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			testCertification := &Certification{
				name:      certification.name,
				year:      certification.year,
				comment:   certification.comment,
				sortIndex: certification.sortIndex,
			}

			err := testCertification.UpdateComment(c.newComment)

			if c.isError {
				assert.Error(t, err)
				assert.Equal(t, certification.comment, testCertification.comment) // 元の値が変わっていないことを確認
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, c.newComment, testCertification.Comment())
		})
	}
}

func TestCertification_UpdateSortIndex(t *testing.T) {
	certification, _ := NewCertification("certification name", makeCertYear("2020-01-01"), "comment", 1)

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
			testCertification := &Certification{
				name:      certification.name,
				year:      certification.year,
				comment:   certification.comment,
				sortIndex: certification.sortIndex,
			}

			err := testCertification.UpdateSortIndex(c.newSortIndex)

			if c.isError {
				assert.Error(t, err)
				assert.Equal(t, certification.sortIndex, testCertification.sortIndex) // 元の値が変わっていないことを確認
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, c.newSortIndex, testCertification.SortIndex())
		})
	}
}