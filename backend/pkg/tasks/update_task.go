package tasks

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateTaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var body UpdateTaskRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.DB.Exec(`UPDATE tasks SET title=$1, description=$2 WHERE id=$3;`, body.Title, body.Description, id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	affRows, _ := result.RowsAffected()

	if affRows == 0 {
		c.AbortWithError(http.StatusNotFound, errors.New("task not found"))
		return
	}

	c.Status(http.StatusOK)
}
