package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(parseID string) (ID, error) {
	id, err := uuid.Parse(parseID)
	return ID(id), err
}
