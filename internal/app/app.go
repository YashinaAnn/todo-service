package app

import (
	"github.com/YashinaAnn/todo-service/internal/database"
	"github.com/YashinaAnn/todo-service/internal/services"
	"github.com/YashinaAnn/todo-service/internal/transport/rest"

	"github.com/gin-gonic/gin"
)

func Run() {
	db, err := database.PrepareDB()
	if err != nil {
		return
	}
	repo := database.NewTodoRepository(db)
	service := services.NewTodoService(repo)
	server := rest.NewServer(service)

	router := gin.Default()
	router.GET("/todos", server.GetTodosHandler)
	router.GET("/todos/:id", server.GetTodoHandler)
	router.POST("/todos", server.AddTodoHandler)
	router.PATCH("/todos/:id", server.ChangeTodoStatusHandler)
	router.Run("localhost:9090")
}
