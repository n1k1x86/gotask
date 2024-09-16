package tasks

import (
	"backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateTaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var body *UpdateTaskRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var task *models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	task.Title = body.Title
	task.Description = body.Description

	h.DB.Save(&task)

	c.JSON(http.StatusOK, &task)

}
