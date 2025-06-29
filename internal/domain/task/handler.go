package task

import (
	"errors"
	"net/http"

	appErrors "todo-list/internal/errors"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}

// helper untuk mengurangi repetisi kode
func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, appErrors.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Resource not found"})
	case errors.Is(err, appErrors.ErrInternal):
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An internal server error occurred"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An unknown error occurred"})
	}
}

func (h *Handler) CreateTask(c *gin.Context) {
	var input CreateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request payload: " + err.Error()})
		return
	}

	task, err := h.Service.CreateTask(input)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Task created successfully", "data": task})
}

func (h *Handler) GetTask(c *gin.Context) {
	tasks, err := h.Service.GetTask()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Tasks fetched successfully", "data": tasks})
}

func (h *Handler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.Service.GetTaskByID(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Task fetched successfully", "data": task})
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var input UpdateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request payload: " + err.Error()})
		return
	}

	task, err := h.Service.UpdateTask(id, input)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Task updated successfully", "data": task})
}

func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteTask(id); err != nil {
		handleError(c, err)
		return // Pastikan ada return setelah menghandle error
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Task deleted successfully"})
}
