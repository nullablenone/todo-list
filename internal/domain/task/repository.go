package task

import "gorm.io/gorm"

type Repository interface {
	Create(task *Task) error
	GetAll() ([]Task, error)
	GetByID(id string) (*Task, error)
	Save(task *Task) error
	Delete(task *Task) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) Create(task *Task) error {
	return r.DB.Create(task).Error
}

func (r *repository) GetAll() ([]Task, error) {
	var tasks []Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *repository) GetByID(id string) (*Task, error) {
	var task Task

	if err := r.DB.Where("ID = ?", id).First(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil

}

func (r *repository) Save(task *Task) error {
	return r.DB.Save(task).Error
}

func (r *repository) Delete(task *Task) error {
	return r.DB.Delete(task).Error
}
