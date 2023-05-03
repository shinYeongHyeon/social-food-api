package shared_core

import "github.com/google/uuid"

func CreateUuid() string {
	return uuid.New().String()
}
