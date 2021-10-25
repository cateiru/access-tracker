package utils

import "github.com/google/uuid"

// Create uniq id.
func CreateId() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", nil
	}
	return id.String(), nil
}
