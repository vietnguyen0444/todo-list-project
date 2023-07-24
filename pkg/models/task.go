package models

import (
	"time"

	"github.com/vietnguyen/todo-list-project/pkg/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// /Note///
// If want to disable auto-increment primary key use: `gorm:"primaryKey;autoIncrement:false"`
// Must be defined json to call APIs
// ///////
type Task struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Order       int32      `json:"order"`
	CreatedAt   time.Time  `json:"createdAt"`
	Status      int8       `json:"status"`
	TaskItems   []TaskItem `json:"taskItems"` // Use for One-To-Many relationship
	//TaskItems []TaskItem `form:many2many:steps` // Use for Many-To-Many relationship
}

// /Note///
// AutoMigrate() will create tables, missing foreign keys, constraints, columns and indexes.
// It will change existing column’s type if its size, precision, nullable changed.
// It WON’T delete unused columns to protect your data.
// ///////
func Init() {
	config.Connect()
	db = config.GetDatabase()
	db.AutoMigrate(&Task{})
}

func (t *Task) CreateTask() (*Task, error) {
	err := db.Create(&t).Error

	return t, err
}

func GetTaskById(id int64) (*Task, error) {
	var task Task
	// db.First(&task, "id = ? ", id)
	// db.First(&task, id)
	err := db.Where("id = ?", id).Find(&task).Error

	return &task, err
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task

	err := db.Find(&tasks).Error

	return tasks, err
}

func (t *Task) UpdateTask() (*Task, error) {
	err := db.Save(&t).Error

	return t, err
}

func DeleteTask(id int64) (Task, error) {
	var task Task

	//db.Delete(&Task{}, id)
	err := db.Where("id = ?", id).Delete(&task).Error

	return task, err
}
