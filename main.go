package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	customGreeting := os.Getenv("CUSTOM_GREETING")
	if customGreeting == "" {
		customGreeting = "Hola!"
	}
	if _, err := fmt.Fprintf(w, "CI CD Testing! %s", customGreeting); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server in port %s\n", port)

	finalPort := fmt.Sprintf(":%s", port)
	if err := http.ListenAndServe(finalPort, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
