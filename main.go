package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	customGreeting := os.Getenv("CUSTOM_GREETING")
	fmt.Println(fmt.Sprintf("CI CD Testing! %s", customGreeting))
}
