package tasks

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRouters(r *gin.Engine, db *gorm.DB) {
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
