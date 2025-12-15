package main

import (
	"fmt"
	"linkedin-automation/internal/browser"
)

func main() {
	fmt.Println("Starting browser...")

	br, _ := browser.StartBrowser()
	defer br.Close()

	fmt.Println("Browser closed. Done.")
}