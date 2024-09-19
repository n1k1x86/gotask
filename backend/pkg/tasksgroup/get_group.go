package tasksgroup

import (
	"github.com/gin-gonic/gin"
)

func (g groupHandler) GetGroup(c *gin.Context) {
	// id, _ := strconv.Atoi(c.Param("id"))

	// var tasks []models.Task
	// var group *models.TaskGroup

	// if result := g.DB.First(&group, id); result.Error != nil {
	// 	c.AbortWithError(http.StatusNotFound, result.Error)
	// 	return
	// }

	// if result := g.DB.Find(&tasks, "group_id = ?", id); result.Error != nil {
	// 	c.AbortWithError(http.StatusBadRequest, result.Error)
	// 	return
	// }

	// group.Tasks = append(group.Tasks, tasks...)

	// c.JSON(http.StatusOK, &group)
}
