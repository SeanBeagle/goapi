package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seanbeagle/goapi/database"
	"github.com/seanbeagle/goapi/endpoint"
)

func main() {
	engine := gin.Default()

	endpoint.RegisterEndpoints(engine)
	database.ConnectDatabase()
	engine.Run(":8080")
}
