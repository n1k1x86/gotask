package tasks

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type handler struct {
	DB *sql.DB
}

func RegisterRouters(r *gin.Engine, db *sql.DB) {
	h := handler{
		DB: db,
	}

	routes := r.Group("/tasks")
	routes.GET("/", h.GetTasks)
	routes.POST("/", h.AddTask)
	routes.PUT("/:id", h.UpdateTask)
	routes.DELETE("/:id", h.DeleteTask)
	routes.GET("/:id", h.GetTask)
}
