package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

// Create uniq id.
func CreateId() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", nil
	}

	hashedValue := sha256.Sum256([]byte(id.String()))

	return hex.EncodeToString(hashedValue[:]), nil
}
