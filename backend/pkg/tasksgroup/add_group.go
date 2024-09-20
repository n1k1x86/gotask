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

	var id uint

	err := g.DB.QueryRow(`INSERT INTO taskgroups (name) VALUES ($1) RETURNING id;`, body.Name).Scan(&id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var group models.TaskGroup = models.TaskGroup{
		ID:   id,
		Name: body.Name,
	}

	c.JSON(http.StatusCreated, &group)
}
