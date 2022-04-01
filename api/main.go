package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seanbeagle/goapi/controller"
	"github.com/seanbeagle/goapi/service"
)

func main() {
	engine := gin.Default()
	service.ConnectDatabase()
	controller.RegisterPersonEndpoints(engine)
	engine.Run(":8080")
}
