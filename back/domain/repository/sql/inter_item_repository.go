package sql

import "devport/domain/model"

type ItemRepository interface {
	Create(item model.Item) error
	Update(item model.Item) model.Item
	Delete(item model.Item) error
}
