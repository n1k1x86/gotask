package tasks

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	result, err := h.DB.Exec(`DELETE FROM tasks WHERE id = $1 RETURNING id`, id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	affectedRows, _ := result.RowsAffected()

	if affectedRows == 0 {
		c.AbortWithError(http.StatusNotFound, errors.New("task not found"))
		return
	}

	c.Status(http.StatusOK)
}
