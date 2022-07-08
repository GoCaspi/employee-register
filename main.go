package main

import (
	"example-project/datasource"
	"example-project/middleware"
	"example-project/model"
	"github.com/gin-gonic/gin"
)

func main() {
	databaseClient := datasource.NewDbClient(model.DbConfig{
		URL:      "mongodb://on4tdb:BMZQuk6pIL39nq46fOQPzygHtrhad5MFQMxs8YBQDW6YsJQSgbsIwO3aeOzlXXEnTjoz7ADVVr9jE1PKzU6GyQ==@on4tdb.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@on4tdb@",
		Database: "office",
	})
	engine := middleware.SetupEngine([]gin.HandlerFunc{middleware.SetupService(databaseClient)})
	engine.Run(":9090")
}
