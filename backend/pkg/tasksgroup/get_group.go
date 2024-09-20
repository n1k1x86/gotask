package tasksgroup

import (
	"backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g groupHandler) GetGroup(c *gin.Context) {
	id := c.Param("id")

	var group models.TaskGroup

	res := g.DB.QueryRow(`SELECT * FROM taskgroups WHERE id = $1`, id)

	err := res.Scan(&group.ID, &group.Name)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &group)
}
