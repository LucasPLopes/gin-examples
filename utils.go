package ginexamples

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	return fmt.Sprintf(":%s", port)
}
