package helpers

import "github.com/google/uuid"

func GenerateUUID() (string, error) {
	uuidv7, err := uuid.NewV7()

	if err != nil {
		return "", err
	}

	return uuidv7.String(), nil
}
