package models

import (
	"time"

	"gorm.io/gorm"
)

//NOTE :
// Conventions
// Primary Key: GORM uses a field named ID as the default primary key for each model.
// Table Names: By default, GORM converts struct names to snake_case and pluralizes them for table names. For instance, a User struct becomes users in the database.
// Column Names: GORM automatically converts struct field names to snake_case for column names in the database.
// Timestamp Fields: GORM uses fields named CreatedAt and UpdatedAt to automatically track the creation and update times of records.

// -> gorm.Model definition
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Updated   int64          `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	Created   int64          `gorm:"autoCreateTime"`      // Use unix seconds as creating time
}

type Person struct {
	// gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
}
