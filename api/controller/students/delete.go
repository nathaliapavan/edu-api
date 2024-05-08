package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api/controller"
	students_usecase "github.com/nathaliapavan/edu-api/usecase/students"
)

func Delete(c *gin.Context) {
	idByParam := c.Params.ByName("id")
	err := students_usecase.Delete(idByParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseErrorMessage("unable to delete student"))
		return
	}

	c.JSON(http.StatusOK, controller.NewResponseSuccessMessage("student removed successfully"))
}
