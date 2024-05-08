package student

import (
	"errors"

	"github.com/nathaliapavan/edu-api/entities"
	"github.com/nathaliapavan/edu-api/shared"
)

func Update(id string, name string, age int) error {
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

	entities.Students[studentIndex].Name = name
	entities.Students[studentIndex].Age = age

	return nil
}
