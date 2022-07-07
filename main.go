package main

import (
	"example-project/datasource"
	"example-project/middleware"
	"example-project/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	databaseClient := datasource.NewDbClient(model.DbConfig{
		URL:       os.Getenv("MONGO_URL"),
		Database:  "office",
	})
	engine := middleware.SetupEngine([]gin.HandlerFunc{middleware.SetupService(databaseClient)})
	engine.Run(":9090")
}
