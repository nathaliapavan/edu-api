package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api/controller"
	students_usecase "github.com/nathaliapavan/edu-api/usecase/students"
)

func Update(c *gin.Context) {
	var input Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseErrorMessage("unable to get payload"))
		return
	}

	idByParam := c.Params.ByName("id")
	if err := students_usecase.Update(idByParam, input.Name, input.Age); err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseErrorMessage("unable to update student"))
		return
	}

	c.JSON(http.StatusOK, controller.NewResponseSuccessMessage("student updated successfully"))
}
