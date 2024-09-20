package tasksgroup

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g groupHandler) DeleteGroup(c *gin.Context) {
	id := c.Param("id")

	res, err := g.DB.Exec(`DELETE FROM taskgroups WHERE id = $1`, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	aR, err := res.RowsAffected()

	if aR == 0 {
		c.AbortWithError(http.StatusNotFound, errors.New("group not found"))
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
