package database

import (
	"database/sql"
	"log"

	"github.com/YashinaAnn/todo-service/internal/models"

	"github.com/go-sql-driver/mysql"
)

func InitDB(config models.Config) (*sql.DB, error) {
	cfg := mysql.Config{
		User:   config.Database.Username,
		Passwd: config.Database.Password,
		Net:    "tcp",
		Addr:   config.Database.Url,
		DBName: config.Database.Name,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
