package model

import (
	gouuid "github.com/satori/go.uuid"
)

type UUID struct {
	id string
}

func NewUUID(id string) UUID {
	if id == "" {
		return UUID{id: gouuid.NewV4().String()}
	} else {
		return UUID{id: id}
	}
}

func (u UUID) ID() string {
	return u.id
}

func IsValidUUID(uuid string) bool {
	_, err := gouuid.FromString(uuid)

	return err == nil
}
