package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seanbeagle/goapi/model"
	"github.com/seanbeagle/goapi/service"
)

func GetPerson(c *gin.Context) {
	var records []model.Person
	service.DB.Find(&records)
	if len(records) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Records"})
		return
	}
	c.JSON(http.StatusOK, records)
}

func GetPersonByID(c *gin.Context) {
	var record model.Person
	if err := service.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}

func CreatePerson(c *gin.Context) {
	// VALIDATE INPUT
	var input model.CreatePersonDTO
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// CREATE RECORD
	record := model.Person{FirstName: input.FirstName, LastName: input.LastName}
	service.DB.Create(&record)
	c.JSON(http.StatusCreated, record)
}

func PatchPerson(c *gin.Context) {
	// GET RECORD
	var record model.Person
	if err := service.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// VALIDATE INPUT
	var input model.UpdatePersonDTO
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.DB.Model(&record).Updates(input)
	c.JSON(http.StatusOK, record)
}

func DeletePerson(c *gin.Context) {
	// GET RECORD
	var record model.Person
	if err := service.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.DB.Delete(&record)
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
