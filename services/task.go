package services

import (
	"errors"
	"todo-app/models"
	"todo-app/repo"
)

type TaskService struct {
	repo *repo.TaskRepository
}

func NewTaskService(repo *repo.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.repo.GetTasks()
}

func (s *TaskService) AddTask(title string) (*models.Task, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	newTask := &models.Task{Title: title, Done: false}
	err := s.repo.AddTask(newTask)
	if err != nil {
		return nil, err
	}
	return newTask, nil
}

func (s *TaskService) CompleteTask(taskID uint) (*models.Task, error) {
	task, err := s.repo.GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	task.Done = true
	err = s.repo.UpdateTask(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) ChangeTaskTitle(taskID uint, newTitle string) error {
	if newTitle == "" {
		return errors.New("title is required")
	}
	return s.repo.ChangeTaskTitle(taskID, newTitle)
}

func (s *TaskService) DeleteTask(taskID uint) error {
	return s.repo.DeleteTask(taskID)
}
