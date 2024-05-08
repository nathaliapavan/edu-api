package infra

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api/controller"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, controller.NewResponseSuccessMessage("ok"))
}
