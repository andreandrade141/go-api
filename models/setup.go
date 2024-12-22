package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/jinzhu/gorm"
// 	"github.com/jinzhu/gorm/dialects/mysql"
// 	"github.com/joho/godotenv"
// )

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error Loading ENV file")
	}

}
