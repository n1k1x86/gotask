package tasksgroup

import (
	"backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g groupHandler) GetGroups(c *gin.Context) {
	var groups []models.TaskGroup = []models.TaskGroup{}

	rows, err := g.DB.Query(`SELECT * FROM taskgroups`)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for rows.Next() {
		var group models.TaskGroup

		err := rows.Scan(&group.ID, &group.Name)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		groups = append(groups, group)
	}

	c.JSON(http.StatusOK, &groups)
}
