package routes

import (
	"todo-list/internal/domain/task"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(taskHandler *task.Handler) *gin.Engine {
	router := gin.Default()
	router.POST("/task", taskHandler.CreateTask)
	router.GET("/task", taskHandler.GetTask)
	router.GET("/task/:id", taskHandler.GetTaskByID)
	router.PUT("/task/:id", taskHandler.UpdateTask)
	router.DELETE("/task/:id", taskHandler.DeleteTask)

	return router
}
