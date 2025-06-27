package task

import "time"

type Task struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"type:varchar(50);not null"`
	Description *string   `gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
