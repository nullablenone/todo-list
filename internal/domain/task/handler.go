package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateTask(c *gin.Context) {
	var input CreateTaskRequest

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request payload",
		})
		return
	}

	task, err := h.Service.CreateTask(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create task",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Success create task",
		"task":    task,
	})
}

func (h *Handler) GetTask(c *gin.Context) {
	tasks, err := h.Service.GetTask()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to get all task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Success get all task",
		"task":    tasks,
	})
}

func (h *Handler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, err := h.Service.GetTaskByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to get task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Success get task",
		"task":    task,
	})

}

func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var input UpdateTaskRequest

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request payload",
		})
		return
	}

	task, err := h.Service.UpdateTask(id, input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to update task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Success update task",
		"task":    task,
	})

}

func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.DeleteTask(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "failed delete task",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success delete task",
	})
}
