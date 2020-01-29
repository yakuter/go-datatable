package inc

import (
	"go-datatable/backend/model"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var DBErr error
var err error

// InitDB opens a database and saves the reference to `Database` struct.
func InitDB() *gorm.DB {

	//driver := dbconf.Database.Driver
	database := "yakuter"
	username := "yakuter"
	password := ""
	host := "localhost"
	port := "5432"

	// POSTGRES
	db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
	if err != nil {
		DBErr = err
		log.Println("db err: ", err)
	}

	// If you want to see the sql queries, uncomment this line
	// db.LogMode(true)

	// If you want to recreate db table, uncomment this line
	// db.DropTableIfExists(&model.Post{})

	db.AutoMigrate(&model.Post{})

	DB = db

	return DB
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

// GetDB helps you to get a connection
func GetDBErr() error {
	return DBErr
}
