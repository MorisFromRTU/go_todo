package repo

import (
	"todo-app/models"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

var taskValidate = validator.New()

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) AddTask(newTask *models.Task) error {
	if err := taskValidate.Struct(newTask); err != nil {
		return err
	}
	return r.DB.Create(newTask).Error
}

func (r *TaskRepository) CompleteTask(taskID uint) (*models.Task, error) {
	// нахождение элемента и его обновление происходят двумя запросами
	var task models.Task
	if err := r.DB.Find(&task, taskID).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Model(&task).Update("done", true).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) ChangeTaskTitle(taskID uint, newTitle string) error {
	return r.DB.Model(&models.Task{}).Where("id = ?", taskID).Update("title", newTitle).Error
}

func (r *TaskRepository) DeleteTask(taskID uint) error {
	return r.DB.Delete(&models.Task{}, taskID).Error
}

func (r *TaskRepository) GetTaskByID(taskid uint) (*models.Task, error) {
	var task *models.Task
	if err := r.DB.Find(&task, taskid).Error; err != nil {
		return nil, err
	}
	return task, nil
}
