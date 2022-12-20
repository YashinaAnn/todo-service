package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/YashinaAnn/todo-service/configs"
	"github.com/YashinaAnn/todo-service/internal/database"
	"github.com/YashinaAnn/todo-service/internal/models"
	"github.com/YashinaAnn/todo-service/internal/services"
	"github.com/YashinaAnn/todo-service/internal/transport/rest"

	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := configs.ReadConfig()
	db := initDB(cfg)
	repo := database.NewTodoRepository(db)
	service := services.NewTodoService(repo)
	server := rest.NewServer(service)

	router := gin.Default()
	router.GET("/todos", server.GetTodosHandler)
	router.GET("/todos/:id", server.GetTodoHandler)
	router.POST("/todos", server.AddTodoHandler)
	router.PATCH("/todos/:id", server.ChangeTodoStatusHandler)
	router.Run(cfg.Server.Url)
}

func initDB(cfg models.Config) *sql.DB {
	db, err := database.InitDB(cfg)
	if err != nil {
		processError(nil)
	}
	return db
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
