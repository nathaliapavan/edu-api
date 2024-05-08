package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api/controller"
	"github.com/nathaliapavan/edu-api/entities"
	"github.com/nathaliapavan/edu-api/shared"
)

func Update(c *gin.Context) {
	var input Input
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadGateway, controller.NewResponseErrorMessage("unable to get payload"))
		return
	}

	var matchStudent entities.Student

	idByParam := c.Params.ByName("id")
	uuid, err := shared.GetUuidByString(idByParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseErrorMessage("error generating uuid"))
		return
	}

	studentIndex := -1
	for i, studentElement := range entities.Students {
		if studentElement.ID == uuid {
			studentIndex = i
			matchStudent = studentElement
			break
		}
	}

	if studentIndex == -1 {
		c.JSON(http.StatusNotFound, controller.NewResponseErrorMessage("student not found"))
		return
	}

	matchStudent.Name = input.Name
	matchStudent.Age = input.Age
	entities.Students[studentIndex] = matchStudent

	c.JSON(http.StatusOK, controller.NewResponseSuccessMessage("student updated successfully"))
}
