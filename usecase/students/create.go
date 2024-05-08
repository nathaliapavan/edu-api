package student

import (
	"github.com/nathaliapavan/edu-api/entities"
)

func Create(name string, age int) (*entities.Student, error) {
	student := entities.NewStudent(name, age)
	entities.Students = append(entities.Students, *student)
	return student, nil
}