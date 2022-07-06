package main

import (
	"example-project/datasource"
	"example-project/middleware"
	"example-project/model"
	"github.com/gin-gonic/gin"
)

func main() {
	databaseClient := datasource.NewDbClient(model.DbConfig{
		URL:      "mongodb://pxdb:F5kUshuDdp8QMHtdc2WuoufKjzpLqoCOb1pyQGtBFi2YQCY7XQtC5B4uKq9se5yk2PJjbgTVCi3hz5y8A16KAA==@pxdb.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&maxIdleTimeMS=120000&appName=@pxdb@",
		Database: "office",
	})
	engine := middleware.SetupEngine([]gin.HandlerFunc{middleware.SetupService(databaseClient)})
	engine.Run(":9090")
}
