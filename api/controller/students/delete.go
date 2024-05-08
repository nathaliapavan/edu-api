package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api/controller"
	"github.com/nathaliapavan/edu-api/entities"
	"github.com/nathaliapavan/edu-api/shared"
)

func Delete(c *gin.Context) {
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
				break
			}
		}

		if studentIndex == -1 {
			c.JSON(http.StatusNotFound, controller.NewResponseErrorMessage("student not found"))
			return
		}

		entities.Students = append(entities.Students[:studentIndex], entities.Students[studentIndex+1:]...)

		c.JSON(http.StatusOK, controller.NewResponseSuccessMessage("student removed successfully"))
}
