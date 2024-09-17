package tasksgroup

import (
	"backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g groupHandler) DeleteGroup(c *gin.Context) {
	id := c.Param("id")

	var group *models.TaskGroup

	if result := g.DB.First(&group, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	g.DB.Delete(&group)

	c.Status(http.StatusOK)
}
