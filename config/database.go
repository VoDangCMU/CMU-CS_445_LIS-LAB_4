package config

import (
	"fmt"
	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	pg_user := "postgres"
	pg_password := "1234567"
	pg_host := "db"
	database_name := "lab_4"

	fmt.Printf("DB connect status: %s:%s@tcp(%s:5432)/%s\n", pg_user, pg_password, pg_host, database_name)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh", pg_host, pg_user, pg_password, database_name)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected")

	err = DB.AutoMigrate(&models.UserInformation{}, &models.UserAuthentication{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return DB
}
