// build
package main

import (
	"log"

	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/config"
	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db := config.InitDB()
	router := routes.SetupRouter(db)
	router.Run(":8080")
}
