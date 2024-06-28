package main

import (
	"gorm/databaseConnect"
	"gorm/models"
)

func init() {
	databaseConnect.ConnectDB()
}

func main() {

	databaseConnect.DB.AutoMigrate(&models.Person{})

	p := models.Person{
		FirstName: "Tanmay",
		LastName:  "Gupta",
		Email:     "tannu@testmail.com",
	}
	databaseConnect.DB.Create(&p)

}
