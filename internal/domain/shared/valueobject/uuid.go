package valueobject

import (
	"github.com/google/uuid"
	"apirest/internal/domain/shared/errors"
)

type UUID struct {
	value uuid.UUID
}

func NewUUID() UUID {
	return UUID{value: uuid.New()}
}

func ParseUUID(s string) (UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return UUID{}, errors.NewInvalidUUIDError(s)
	}
	return UUID{value: id}, nil
}

func MustParseUUID(s string) UUID {
	id, err := ParseUUID(s)
	if err != nil {
		panic(err)
	}
	return id
}

func (u UUID) String() string {
	return u.value.String()
}

func (u UUID) Value() uuid.UUID {
	return u.value
}

func (u UUID) IsZero() bool {
	return u.value == uuid.Nil
}

func (u UUID) Equals(other UUID) bool {
	return u.value == other.value
}
