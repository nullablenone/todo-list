package task

import "fmt"

type Service interface {
	CreateTask(input CreateTaskRequest) (*Task, error)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) CreateTask(input CreateTaskRequest) (*Task, error) {
	task := Task{
		Title:       input.Title,
		Description: input.Description,
	}

	if err := s.Repo.Create(&task); err != nil {
		return nil, fmt.Errorf("ailed to create task: %w", err)
	}

	return &task, nil

}
