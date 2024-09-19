package tasks

import (
	"backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	row := h.DB.QueryRow(`SELECT * FROM tasks WHERE id = $1`, id)

	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.GroupID); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, &task)
}
