package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api/controller"
	students_usecase "github.com/nathaliapavan/edu-api/usecase/students"
)

func Details(c *gin.Context) {
	idByParam := c.Params.ByName("id")
	student, err := students_usecase.Details(idByParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseErrorMessage("error getting student details"))
		return
	}

	if student == nil {
		c.JSON(http.StatusNotFound, controller.NewResponseErrorMessage("student not found"))
		return
	}

	c.JSON(http.StatusOK, student)
}
