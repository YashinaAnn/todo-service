package services

import (
	"github.com/YashinaAnn/todo-service/internal/database"
	"github.com/YashinaAnn/todo-service/internal/models"
)

type TodoService struct {
	repository *database.TodoRepository
}

func NewTodoService(repo *database.TodoRepository) *TodoService {
	return &TodoService{repository: repo}
}

func (service *TodoService) GetTodoById(id int64) (models.Todo, error) {
	return service.repository.GetTodoById(id)
}

func (service *TodoService) GetTodos() ([]models.Todo, error) {
	return service.repository.GetAllTodos()
}

func (service *TodoService) AddTodo(t models.Todo) (models.Todo, error) {
	id, err := service.repository.SaveTodo(t)
	if err != nil {
		return models.Todo{}, err
	}
	t.ID = id
	return t, nil
}
