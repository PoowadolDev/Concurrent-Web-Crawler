package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting Concurrent Web Crawler")

	content, err := fetch("https://www.w3schools.com/")
	if err != nil {
		fmt.Println("Error fetching pages:", err)
	}

	fmt.Println("Page Contents")
	fmt.Println(content)
}
