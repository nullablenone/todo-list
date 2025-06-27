package task

import "gorm.io/gorm"

type Repository interface {
	Create(task *Task) error
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
