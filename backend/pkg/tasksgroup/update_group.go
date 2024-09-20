package tasksgroup

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupUpdateRequestBody struct {
	Name string `json:"name"`
}

func (g groupHandler) UpdateGroup(c *gin.Context) {
	id := c.Param("id")

	var body GroupAddRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := g.DB.Exec(`UPDATE taskgroups SET name = $1 WHERE id = $2`, body.Name, id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	aR, err := res.RowsAffected()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if aR == 0 {
		c.AbortWithError(http.StatusNotFound, errors.New("group not found"))
		return
	}

	c.Status(http.StatusOK)
}
