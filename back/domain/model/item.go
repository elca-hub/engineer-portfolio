package model

import (
	"errors"
	"fmt"
	"time"
)

type (
	Item struct {
		id        int64
		title     string
		memo      string
		folder    Folder
		isIdea    bool
		createdAt time.Time
		updatedAt time.Time
		parents   []Item // コネクト元
		children  []Item // コネクト先
	}
)

const (
	maxItemNameLength = 100
	maxItemMemoLength = 1000
)

func NewItem(
	id int64,
	name string,
	memo string,
	isIdea bool,
	folder *Folder,
	createdAt time.Time,
	updatedAt time.Time,
	parents []Item,
	children []Item,
) (Item, error) {
	if len(name) > maxItemNameLength {
		return Item{}, errors.New(fmt.Sprintf("Item name is too long: %s", name))
	}
	if len(name) == 0 {
		return Item{}, errors.New("item name is empty")
	}

	if len(memo) > maxItemMemoLength {
		return Item{}, errors.New(fmt.Sprintf("Item memo is too long: %s", memo))
	}

	return Item{
		id:        id,
		title:     name,
		memo:      memo,
		isIdea:    isIdea,
		folder:    *folder,
		createdAt: createdAt,
		updatedAt: updatedAt,
		parents:   parents,
		children:  children,
	}, nil
}

func (i Item) ID() int64 {
	return i.id
}

func (i Item) Name() string {
	return i.title
}

func (i Item) Memo() string {
	return i.memo
}

func (i Item) IsIdea() bool {
	return i.isIdea
}

func (i Item) Folder() *Folder {
	return &i.folder
}

func (i Item) CreatedAt() time.Time {
	return i.createdAt
}

func (i Item) UpdatedAt() time.Time {
	return i.updatedAt
}
func (i Item) Parents() []Item {
	return i.parents
}

func (i Item) Children() []Item {
	return i.children
}
