package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seanbeagle/goapi/database"
	"github.com/seanbeagle/goapi/model"
)

func GetPerson(c *gin.Context) {
	var people []model.Person
	database.DB.Find(&people)
	if len(people) == 0 {
		c.String(http.StatusNotFound, "No people")
		return
	}
	c.JSON(http.StatusOK, people)
}

func GetPersonByID(c *gin.Context) {
	var person model.Person
	if err := database.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
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
	database.DB.Create(&person)
	c.JSON(http.StatusOK, person)
}

func PatchPerson(c *gin.Context) {
	// GET PERSON
	var person model.Person
	if err := database.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// VALIDATE INPUT
	var input model.UpdatePerson
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&person).Updates(input)
	c.JSON(http.StatusOK, person)
}

func DeletePerson(c *gin.Context) {
	// Get model if exist
	var person model.Person
	if err := database.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	database.DB.Delete(&person)
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
