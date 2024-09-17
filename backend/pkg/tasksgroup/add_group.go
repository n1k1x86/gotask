package tasksgroup

import (
	"backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupAddRequestBody struct {
	Name string `json:"name"`
}

func (g groupHandler) AddGroup(c *gin.Context) {
	var body *GroupAddRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var group *models.TaskGroup

	group.Name = body.Name

	if result := g.DB.Create(&group); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, &group)

}
