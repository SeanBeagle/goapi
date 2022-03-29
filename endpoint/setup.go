package endpoint

import (
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.GET("/person", GetPerson)
		v1.GET("/person/:id", GetPersonByID)
		v1.POST("/person", CreatePerson)
		v1.DELETE("/person/:id", DeletePerson)
		v1.PATCH("/person/:id", PatchPerson)
	}
}
