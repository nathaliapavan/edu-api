package student

import (
	"github.com/google/uuid"
	"github.com/nathaliapavan/edu-api/entities"
	"github.com/nathaliapavan/edu-api/shared"
)

func Details(id string) (*entities.Student, error) {
	uuid, err := shared.GetUuidByString(id)
	if err != nil {
		return nil, err
	}

	student, err := findStudentByID(uuid)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func findStudentByID(uuid uuid.UUID) (*entities.Student, error) {
	for _, student := range entities.Students {
		if student.ID == uuid {
			return &student, nil
		}
	}
	return nil, nil
}
