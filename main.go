package main

import (
	"fmt"
	"gorm/databaseConnect"
	loadenvvariables "gorm/loadEnvVariables"
	"gorm/models"

	"github.com/gin-gonic/gin"
)

func init() {
	loadenvvariables.LoadEnvVariables()
	databaseConnect.ConnectDB()
}

func main() {

	databaseConnect.DB.AutoMigrate(&models.Person{})

	router := gin.Default()
	fmt.Println("Server is getting started...")

	//ROUTES
	router.POST("api/person", databaseConnect.CreatePerson)
	router.GET("api/persons", databaseConnect.GetPerson)
	router.DELETE("api/deleteperson/:first_name", databaseConnect.DeletePersonByName)
	router.GET("api/persons/:first_name", databaseConnect.GetPersonByName)
	// router.PUT("api/persons/:first_name", databaseConnect.UpdatePersonByName)

	router.Run()
}
