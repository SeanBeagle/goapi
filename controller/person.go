package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seanbeagle/goapi/model"
	"github.com/seanbeagle/goapi/service"
)

func GetPerson(c *gin.Context) {
	var people []model.Person
	service.DB.Find(&people)
	if len(people) == 0 {
		c.String(http.StatusNotFound, "No people")
		return
	}
	c.JSON(http.StatusOK, people)
}

func GetPersonByID(c *gin.Context) {
	var person model.Person
	if err := service.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func CreatePerson(c *gin.Context) {
	// VALIDATE INPUT
	var input model.CreatePerson
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// CREATE PERSON
	person := model.Person{FirstName: input.FirstName, LastName: input.LastName}
	service.DB.Create(&person)
	c.JSON(http.StatusOK, person)
}

func PatchPerson(c *gin.Context) {
	// GET PERSON
	var person model.Person
	if err := service.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// VALIDATE INPUT
	var input model.UpdatePerson
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.DB.Model(&person).Updates(input)
	c.JSON(http.StatusOK, person)
}

func DeletePerson(c *gin.Context) {
	// Get model if exist
	var person model.Person
	if err := service.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	service.DB.Delete(&person)
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func RegisterPersonEndpoints(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.GET("/person", GetPerson)
		v1.GET("/person/:id", GetPersonByID)
		v1.POST("/person", CreatePerson)
		v1.DELETE("/person/:id", DeletePerson)
		v1.PATCH("/person/:id", PatchPerson)
	}
}
