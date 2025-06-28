package task

import "fmt"

type Service interface {
	CreateTask(input CreateTaskRequest) (*Task, error)
	GetTask() ([]Task, error)
	GetTaskByID(id string) (*Task, error)
	UpdateTask(id string, input UpdateTaskRequest) (*Task, error)
	DeleteTask(id string) error
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
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return &task, nil

}

func (s *service) GetTask() ([]Task, error) {
	tasks, err := s.Repo.GetAll()

	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %w", err)
	}

	return tasks, nil
}

func (s *service) GetTaskByID(id string) (*Task, error) {

	task, err := s.Repo.GetByID(id)

	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	return task, nil

}

func (s *service) UpdateTask(id string, input UpdateTaskRequest) (*Task, error) {

	task, err := s.Repo.GetByID(id)

	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	task.Title = input.Title
	task.Description = input.Description

	err = s.Repo.Save(task)
	if err != nil {
		return nil, fmt.Errorf("failed to save tasks: %w", err)
	}

	return task, nil
}

func (s *service) DeleteTask(id string) error {
	task, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("failed to get tasks: %w", err)
	}

	err = s.Repo.Delete(task)
	if err != nil {
		return fmt.Errorf("failed to delete tasks: %w", err)
	}

	return nil
}
