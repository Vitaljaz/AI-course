package main

import (
	"fmt"
	"time"
)

var (
	N int
	startX int
	startY int
	endX int
	endY int
)

func RunBruteForce() {
	timer := time.Now()
	BruteForce(N, startX, startY, endX, endY)
	fmt.Printf("[BruteForce]: Elapsed time: %v", time.Since(timer))
}


func main() {
	fmt.Println("Enter N:")
	fmt.Scan(&N)

	fmt.Println("Enter start position (x;y): ")
	fmt.Scan(&startX, &startY)

	fmt.Println("Enter end position (x;y): ")
	fmt.Scan(&endX, &endY)
}