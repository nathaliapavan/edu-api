package entities

import (
	"github.com/google/uuid"
	"github.com/nathaliapavan/edu-api/shared"
)

type Student struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

func NewStudent(name string, age int) *Student {
	return &Student{
		ID:   shared.GetUuid(),
		Name: name,
		Age:  age,
	}
}

var Students = []Student{
	{ID: shared.GetUuid(), Name: "Peter Parker", Age: 18},
}
