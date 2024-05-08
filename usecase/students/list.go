package student

import "github.com/nathaliapavan/edu-api/entities"

func List() ([]entities.Student, error) {
	return entities.Students, nil
}