package tasksgroup

import (
	"backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g groupHandler) GetGroups(c *gin.Context) {
	var groups []*models.TaskGroup

	if result := g.DB.Find(&groups); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, &groups)
}
