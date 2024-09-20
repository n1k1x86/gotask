package tasks

import (
	"backend/pkg/common/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	row := h.DB.QueryRow(`SELECT * FROM tasks WHERE id = $1`, id)
	var dueDate sql.NullTime

	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.GroupID, &dueDate); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	if dueDate.Valid {
		task.DueDate = dueDate.Time.Format("2006-01-02")
	}

	c.JSON(http.StatusOK, &task)
}
