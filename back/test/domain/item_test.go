package domain

import (
	"github.com/stretchr/testify/assert"
	"devport/domain/model"
	"testing"
	"time"
)

func TestItem(t *testing.T) {
	folder := model.NewFolder(1, "testFolder")
	_, err := model.NewItem(1, "test", "", true, folder, time.Now(), time.Now(), []model.Item{}, []model.Item{})

	assert.NoError(t, err)
}

func TestMinNameLen(t *testing.T) {
	folder := model.NewFolder(1, "testFolder")
	_, err := model.NewItem(1, "", "", true, folder, time.Now(), time.Now(), []model.Item{}, []model.Item{})

	assert.Error(t, err)

	_, err = model.NewItem(1, "a", "", true, folder, time.Now(), time.Now(), []model.Item{}, []model.Item{})

	assert.NoError(t, err)
}

func TestMaxNameLen(t *testing.T) {
	// 100文字
	maxLen := 100
	var checker string

	for i := 0; i < maxLen; i++ {
		checker += "a"
	}

	folder := model.NewFolder(1, "testFolder")
	_, err := model.NewItem(1, checker, "", true, folder, time.Now(), time.Now(), []model.Item{}, []model.Item{})

	assert.NoError(t, err)

	// 101文字
	_, err = model.NewItem(1, checker+"a", "", true, folder, time.Now(), time.Now(), []model.Item{}, []model.Item{})

	assert.Error(t, err)
}

func TestMaxMemoLen(t *testing.T) {
	maxLen := 1000
	var checker string

	for i := 0; i < maxLen; i++ {
		checker += "a"
	}

	folder := model.NewFolder(1, "testFolder")
	_, err := model.NewItem(1, "test", checker, true, folder, time.Now(), time.Now(), []model.Item{}, []model.Item{})

	assert.NoError(t, err)

	// 1001文字
	_, err = model.NewItem(1, "test", checker+"a", true, folder, time.Now(), time.Now(), []model.Item{}, []model.Item{})

	assert.Error(t, err)
}
