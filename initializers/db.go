package initializers

import (
	"hditano/MVCProject/models"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initialize gorm
var DB *gorm.DB

func ConnectToDatabase() {
	os.Remove("./database.db")
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	checkErrors(err)

}

func checkErrors(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Customer{})
}
