package tasks

import (
	"fmt"
	"net/http"
	"time"

	"backend/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type TaskAddRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	GroupId     int    `json:"groupId"`
	DueDate     string `json:"dueDate"`
}

func parseTime(body TaskAddRequestBody) (time.Time, error) {
	dueDate, err := time.Parse("2006-01-02", body.DueDate)
	if err != nil {
		return time.Time{}, err
	}
	return dueDate, nil
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

	dueDate, err := parseTime(body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.DB.Exec(`INSERT INTO tasks(title, description, groupid, duedate) VALUES ($1, $2, $3, $4) RETURNING id;`, body.Title, body.Description, body.GroupId, dueDate)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(result)

	c.Status(http.StatusCreated)
}
