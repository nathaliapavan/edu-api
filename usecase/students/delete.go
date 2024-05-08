package student

import (
	"errors"

	"github.com/nathaliapavan/edu-api/entities"
	"github.com/nathaliapavan/edu-api/shared"
)

func Delete(id string) error {
	uuid, err := shared.GetUuidByString(id)
	if err != nil {
		return err
	}

	studentIndex := -1
	for i, student := range entities.Students {
		if student.ID == uuid {
			studentIndex = i
			break
		}
	}

	if studentIndex == -1 {
		return errors.New("student not found")
	}

	entities.Students = append(entities.Students[:studentIndex], entities.Students[studentIndex+1:]...)
	return nil
}