package main

import (
	"fmt"
	"gorm/databaseConnect"
	"gorm/models"

	"github.com/gin-gonic/gin"
)

func init() {
	databaseConnect.ConnectDB()
}

func main() {

	databaseConnect.DB.AutoMigrate(&models.Person{})

	// Create a new user
	// create.CreatePersen(databaseConnect.DB, "Alice", "Singh", "alice@example.com")

	// Retrieve a user by ID
	// get.GetPerson(databaseConnect.DB, "Alice")

	//Delete a user
	// deleting.DeletePersonByID(databaseConnect.DB, "Raju")

	router := gin.Default()
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 8080 ...")

	//ROUTES
	router.POST("api/person", databaseConnect.CreatePerson)
	router.GET("api/persons", databaseConnect.GetPerson)
	router.DELETE("api/deleteperson/:first_name", databaseConnect.DeletePersonByName)
	router.GET("api/persons/:first_name", databaseConnect.GetPersonByName)

	router.Run(":8080")
}
