package tasksgroup

import (
	"backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupUpdateRequestBody struct {
	Name string `json:"name"`
}

func (g groupHandler) UpdateGroup(c *gin.Context) {
	// id := c.Param("id")

	var body *GroupUpdateRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var group *models.TaskGroup

	// if result := g.DB.First(&group, id); result.Error != nil {
	// 	c.AbortWithError(http.StatusNotFound, result.Error)

	// }

	group.Name = body.Name

	// g.DB.Save(&group)

	c.JSON(http.StatusOK, &group)
}
