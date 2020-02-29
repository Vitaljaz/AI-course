package main

import (
	"fmt"
)

func main() {
	var N int
	var startX, startY int
	var endX, endY int

	fmt.Println("Enter N:")
	fmt.Scan(&N)

	fmt.Println("Enter start position (x;y): ")
	fmt.Scan(&startX, &startY)

	fmt.Println("Enter end position (x;y): ")
	fmt.Scan(&endX, &endY)
}