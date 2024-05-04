package build

import (
	"app/internal/api/router"
	"log"

	"github.com/joho/godotenv"
)

func Build() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println()
		return
	}
	app := router.Setup()
	app.Run(":8080")
}
