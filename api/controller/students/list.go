package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/entities"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, entities.Students)
}
