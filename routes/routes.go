package routes

import (
	"todo-list/internal/domain/task"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(taskHandler *task.Handler) *gin.Engine {
	router := gin.Default()
	router.POST("/task", taskHandler.CreateTask)

	return router
}
