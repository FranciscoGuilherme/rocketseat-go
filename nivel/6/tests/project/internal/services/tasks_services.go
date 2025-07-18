package services

import "rocketseat-go/internal/store"

type TaskService struct {
	Store store.TaskStore
}

func NewTaskService(store store.TaskStore) *TaskService {
	return &TaskService{Store: store}
}

func (s *TaskService) CreateTask(
	title,
	description string,
	priority int32,
) (store.Task, error) {
	task, err := s.Store.CreateTask(title, description, priority)
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskService) GetTaskByID(id int32) (store.Task, error) {
	task, err := s.Store.GetTaskByID(id)
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskService) ListTasks() ([]store.Task, error) {
	tasks, err := s.Store.ListTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) UpdateTask(
	id int32,
	title, description string,
	priority int32,
) (store.Task, error) {
	task, err := s.Store.UpdateTask(id, title, description, priority)
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id int32) error {
	err := s.Store.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
