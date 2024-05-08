package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api/controller"
	students_usecase "github.com/nathaliapavan/edu-api/usecase/students"
)

func Create(c *gin.Context) {
	var input Input
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseErrorMessage("unable to get payload"))
		return
	}

	student, err := students_usecase.Create(input.Name, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseErrorMessage("unable to create student"))
		return
	}

	c.JSON(http.StatusCreated, student)
}
