// build
package main

import (
	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/config"
	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/routes"
)

func main() {
	db := config.InitDB()
	router := routes.SetupRouter(db)
	router.Run(":8080")
}
