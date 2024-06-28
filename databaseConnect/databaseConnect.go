package databaseConnect

import (
	"fmt"
	"gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:2003@society@/societydb"), &gorm.Config{})

	if err != nil {
		panic(" Could not connect mysql DB ")
	}

	fmt.Println("Database connection is successful !!")
}

// CREATING A RECORD
func CreatePerson(c *gin.Context) {

	var p models.Person

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := DB.Create(&p).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}

// FETCHING ALL THE RECORDS
func GetPerson(c *gin.Context) {

	var users []models.Person

	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// DELETING A RECORD
func DeletePersonByName(c *gin.Context) {

	name := c.Param("first_name")
	if err := DB.Where("first_name = ?", name).Delete(&models.Person{}).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func GetPersonByName(c *gin.Context) {

	name := c.Param("first_name")
	var user models.Person
	if err := DB.Where("first_name = ?", name).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}
