package tasks

import (
	"backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetTasks(c *gin.Context) {
	var tasks []models.Task = []models.Task{}

	rows, err := h.DB.Query("SELECT * FROM tasks;")

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.GroupID)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, &tasks)
}
