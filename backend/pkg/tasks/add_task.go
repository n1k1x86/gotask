package tasks

import (
	"net/http"

	"backend/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type TaskAddRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h handler) AddTask(c *gin.Context) {
	body := TaskAddRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	task.Title = body.Title
	task.Description = body.Description

	if result := h.DB.Create(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &task)
}
