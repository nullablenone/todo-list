package main

import (
	"log"
	"todo-list/config"
	"todo-list/internal/domain/task"
	"todo-list/routes"
)

func main() {
	config.LoadENV()
	db := config.ConnectDB()

	err := db.AutoMigrate(task.Task{})
	if err != nil {
		log.Fatalf("error failed to create table: %v", err)
	}

	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := task.NewHandler(taskService)

	router := routes.SetupRoutes(taskHandler)
	router.Run()
}
