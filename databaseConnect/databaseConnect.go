package databaseConnect

import (
	"fmt"

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
