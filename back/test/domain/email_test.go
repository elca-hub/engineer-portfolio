package domain

import (
	"github.com/stretchr/testify/assert"
	"devport/domain/model"
	"testing"
)

func TestTrueEmail(t *testing.T) {
	_, err := model.NewEmail("test@example.com")

	assert.NoError(t, err)
}

func TestFalseEmail(t *testing.T) {
	_, err := model.NewEmail("testexample.com")

	assert.Error(t, err)
}
