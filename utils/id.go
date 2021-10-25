package utils

import "github.com/google/uuid"

func CreateId() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", nil
	}
	return id.String(), nil
}
