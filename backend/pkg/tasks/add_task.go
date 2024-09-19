package tasks

import (
	"fmt"
	"net/http"

	"backend/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type TaskAddRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	GroupId     int    `json:"groupId"`
}

func (h handler) AddTask(c *gin.Context) {
	body := TaskAddRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var group models.TaskGroup
	row := h.DB.QueryRow(`SELECT * FROM taskgroups WHERE id = $1`, body.GroupId)

	if err := row.Scan(&group.ID, &group.Name); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	result, err := h.DB.Exec(`INSERT INTO tasks(title, description, groupid) VALUES ($1, $2, $3) RETURNING id;`, body.Title, body.Description, body.GroupId)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(result)

	c.Status(http.StatusCreated)
}
