package domain

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

func ParseIDOrNil(s string) ID {
	id, _ := ParseID(s)
	return id
}
