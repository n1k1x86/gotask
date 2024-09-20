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

	var group models.TaskGroup = models.TaskGroup{
		Name: body.Name,
	}

	_, err := g.DB.Exec(`INSERT INTO taskgroups (name) VALUES ($1)`, group.Name)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, &group)
}
