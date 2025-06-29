package task

type CreateTaskRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description,omitempty"`
}

type UpdateTaskRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}
