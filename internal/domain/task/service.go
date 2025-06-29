package task

import (
	"errors"
	"log"
	appErrors "todo-list/internal/errors"

	"gorm.io/gorm"
)

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
		log.Printf("ERROR: Repository Create failed: %v", err)
		return nil, appErrors.ErrInternal
	}

	return &task, nil

}

func (s *service) GetTask() ([]Task, error) {
	tasks, err := s.Repo.GetAll()

	if err != nil {
		log.Printf("ERROR: Repository GetAll failed: %v", err)
		return nil, appErrors.ErrInternal
	}

	return tasks, nil
}

func (s *service) GetTaskByID(id string) (*Task, error) {

	task, err := s.Repo.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrNotFound
		}
		log.Printf("ERROR: Repository GetByID failed for id %s: %v", id, err)
		return nil, appErrors.ErrInternal
	}

	return task, nil

}

func (s *service) UpdateTask(id string, input UpdateTaskRequest) (*Task, error) {

	task, err := s.Repo.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrNotFound
		}
		log.Printf("ERROR: Repository GetByID failed for id %s: %v", id, err)
		return nil, appErrors.ErrInternal
	}

	if input.Title != "" {
		task.Title = input.Title
	}

	if input.Description != nil {
		task.Description = input.Description
	}

	err = s.Repo.Save(task)
	if err != nil {
		log.Printf("ERROR: Repository Save failed for id %s: %v", id, err)
		return nil, appErrors.ErrInternal
	}

	return task, nil
}

func (s *service) DeleteTask(id string) error {
	task, err := s.Repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return appErrors.ErrNotFound
		}
		log.Printf("ERROR: Repository GetByID failed for id %s: %v", id, err)
		return appErrors.ErrInternal
	}

	err = s.Repo.Delete(task)
	if err != nil {
		log.Printf("ERROR: Repository Delete failed for id %d: %v", task.ID, err)
		return appErrors.ErrInternal
	}

	return nil
}
