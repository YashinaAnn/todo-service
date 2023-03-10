package database

import (
	"database/sql"
	"fmt"

	"github.com/YashinaAnn/todo-service/internal/models"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (repo *TodoRepository) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo

	rows, err := repo.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, fmt.Errorf("GetAllTodos: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Item, &todo.Completed); err != nil {
			return nil, fmt.Errorf("GetAllTodos: %v", err)
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllTodos: %v", err)
	}

	return todos, nil
}

func (repo *TodoRepository) GetTodoById(id int64) (models.Todo, error) {
	var todo models.Todo
	row := repo.db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&todo.ID, &todo.Item, &todo.Completed); err != nil {
		if err == sql.ErrNoRows {
			return todo, fmt.Errorf("GetTodoById %d: no such todo", id)
		}
		return todo, fmt.Errorf("GetTodoById %d: %v", id, err)
	}
	return todo, nil
}

func (repo *TodoRepository) SaveTodo(todo models.Todo) (int64, error) {
	result, err := repo.db.Exec("INSERT INTO todos (ID, item, completed) VALUES (?, ?, ?)", todo.ID, todo.Item, todo.Completed)
	if err != nil {
		return 0, fmt.Errorf("SaveTodo: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("SaveTodo: %v", err)
	}
	return id, nil
}
