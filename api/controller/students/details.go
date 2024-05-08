package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api/controller"
	"github.com/nathaliapavan/edu-api/entities"
	"github.com/nathaliapavan/edu-api/shared"
)

func Details(c *gin.Context) {
		idByParam := c.Params.ByName("id")
		uuid, err := shared.GetUuidByString(idByParam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, controller.NewResponseErrorMessage("error generating uuid"))
			return
		}

		studentIndex := -1
		var studentFound entities.Student
		for i, studentElement := range entities.Students {
			if studentElement.ID == uuid {
				studentIndex = i
				studentFound = studentElement
				break
			}
		}

		if studentIndex == -1 {
			c.JSON(http.StatusNotFound, controller.NewResponseErrorMessage("student not found"))
			return
		}

	c.JSON(http.StatusOK, studentFound)
}
