package rest

import (
	"net/http"
	"strconv"

	"github.com/YashinaAnn/todo-service/internal/models"
	"github.com/YashinaAnn/todo-service/internal/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	service *services.TodoService
}

func NewServer(s *services.TodoService) *Server {
	return &Server{service: s}
}

func (server *Server) GetTodosHandler(context *gin.Context) {
	todos, err := server.service.GetTodos()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	context.IndentedJSON(http.StatusOK, todos)
}

func (server *Server) AddTodoHandler(context *gin.Context) {
	var newTodo models.Todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	newTodo, err := server.service.AddTodo(newTodo)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func (server *Server) GetTodoHandler(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	todo, err := server.service.GetTodoById(int64(id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func (server *Server) ChangeTodoStatusHandler(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	t, err := server.service.GetTodoById(int64(id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	t.Completed = !t.Completed
	context.IndentedJSON(http.StatusOK, t)
}
