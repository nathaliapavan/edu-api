package shared

import (
	"log"

	"github.com/google/uuid"
)

func GetUuid() uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Error", err)
	}
	return uuid
}

func GetUuidByString(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}
