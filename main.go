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

	links, err := parseLinks(content)
	if err != nil {
		fmt.Println("Error parsing pages:", err)
	}

	fmt.Println("Found links:")
	for _, link := range links {
		fmt.Println(link)
	}
}
