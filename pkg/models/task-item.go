package models

import (
	"gorm.io/gorm"
)

type TaskItem struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       int32  `json:"order"`
	TaskID      uint   `json:"task_id"`
}

func (ti *TaskItem) CreateTaskItem() (*TaskItem, error) {
	err := db.Create(&ti).Error
	return ti, err
}
