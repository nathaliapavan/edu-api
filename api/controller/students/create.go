package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api/controller"
	"github.com/nathaliapavan/edu-api/entities"
)

func Create(c *gin.Context) {
	var input Input
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadGateway, controller.NewResponseErrorMessage("unable to get payload"))
		return
	}

	student := entities.NewStudent(input.Name, input.Age)
	entities.Students = append(entities.Students, *student)

	c.JSON(http.StatusCreated, student)
}
