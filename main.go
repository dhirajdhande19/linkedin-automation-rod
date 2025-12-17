package main

import (
	"fmt"

	"linkedin-automation/internal/browser"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using system env")
	}

	fmt.Println("Starting browser...")

	br, _ := browser.StartBrowser()
	defer br.Close()

	fmt.Println("Browser closed. Done.")
}