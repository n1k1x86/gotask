package tasksgroup

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type groupHandler struct {
	DB *sql.DB
}

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
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
