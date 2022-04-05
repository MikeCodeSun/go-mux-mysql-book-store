package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	// get .env varible
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("Error dotenv")
	}
	username := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_PASSWORD")
	//  fomat connet addrres
	Dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/book_store?charset=utf8mb4&parseTime=True&loc=Local", username, password)
	var errDb error
	db, errDb = gorm.Open(mysql.Open(Dsn))
	if errDb != nil {
		
		log.Fatal(errDb)
	}
}

func GetDB() *gorm.DB{
	return db
}

