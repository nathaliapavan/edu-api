package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api/controller"
	students_usecase "github.com/nathaliapavan/edu-api/usecase/students"
)

func List(c *gin.Context) {
	students, err := students_usecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseErrorMessage("unable to list students"))
		return
	}

	c.JSON(http.StatusOK, students)
}
