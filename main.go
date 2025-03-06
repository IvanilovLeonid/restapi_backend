package main

import (
	"hw1/api/http"
	"hw1/repository/dbram"
	"hw1/usecases/service"
	"log"

	_ "hw1/docs" // Загрузка сгенерированных файлов Swagger
)

// @title Task Manager API
// @version 1.0
// @description This is a simple API for managing tasks with status and result tracking.
// @host localhost:8080
// @BasePath /
func main() {
	port := ":8080"

	dbRam := dbram.NewObject()
	db := service.NewObject(dbRam)
	log.Println("server is running on port " + port)

	err := http.CreateAndRunServer(db, port)

	if err != nil {
		log.Fatalf("Fail to start %v\n", err)
	}
}
