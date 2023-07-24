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

func GetTaskItemById(id int64) (*TaskItem, error) {
	var taskItem TaskItem

	err := db.Where("id = ?", id).Find(&taskItem).Error

	return &taskItem, err
}

func GetAllTaskItems() ([]TaskItem, error) {
	var taskItems []TaskItem

	err := db.Find(&taskItems).Error

	return taskItems, err
}

func (ti *TaskItem) UpdateTaskItem() (*TaskItem, error) {
	err := db.Save(&ti).Error

	return ti, err
}

func DeleteTaskItem(id int64) (TaskItem, error) {
	var taskItem TaskItem
	err := db.Where("id = ?", id).Delete(&taskItem).Error

	return taskItem, err
}
