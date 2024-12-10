package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=aws-0-ap-southeast-1.pooler.supabase.com user=postgres.auodbtvdrjrpblxxxxvn password=!Bombompokemon3000 dbname=postgres port=6543 sslmode=require TimeZone=Asia/Jakarta"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Database connected successfully.")
}

func GetDB() *gorm.DB {
	return DB
}
