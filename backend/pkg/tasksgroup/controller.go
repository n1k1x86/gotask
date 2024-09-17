package tasksgroup

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type groupHandler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	g := groupHandler{
		DB: db,
	}

	routes := r.Group("/taskgroups")

	routes.GET("/", g.GetGroups)
	routes.GET("/:id", g.GetGroup)
	routes.PUT("/:id", g.UpdateGroup)
	routes.DELETE("/:id", g.DeleteGroup)
	routes.POST("/", g.AddGroup)
}
