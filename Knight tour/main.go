package main

import (
	"fmt"
	"time"
	"os"
)

var (
	n 	   int
	startX int
	startY int
	endX   int
	endY   int
	destroyCount     int
	destroyFieldsX []int
	destroyFieldsY []int
)

func RunBruteForce() {
	fmt.Println("[BruteForce]:")
	timer := time.Now()
	var current int
	chess := CreateChess()
	if BruteForce(chess, current, startX, startY) {
		PrintChess(chess)
		fmt.Printf("[BruteForce]: Elapsed time: %v", time.Since(timer))
	} else {
		fmt.Println("[BruteForce]: Impossible!")
	}
}

func PrintChess(chess [][]int) {
	chess[startX][startY] = -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%3d", chess[i][j])
		}
		fmt.Println()
	}
}

func CreateChess() [][] int {
	chess := make([][]int, n)
	for i := range chess {
		chess[i] = make([]int, n)
	}

	for i := 0; i < destroyCount; i++ {
		chess[destroyFieldsX[i]][destroyFieldsY[i]] = -3
	}
	return chess
}

func CheckingArguments() {
	if n <= 1 { 
		fmt.Println("[Error]: Bad N value")
		os.Exit(3)
	}

	if startX >= n || startY >= n || startX < 0 || startY < 0 {
		fmt.Println("[Error]: Bad start values")
		os.Exit(3)
	}

	if endX >= n || endY >= n || endX < 0 || endY < 0 {
		fmt.Println("[Error]: Bad end values")
		os.Exit(3)
	}

	if destroyCount < 0 {
		fmt.Println("[Error]: Bad destroyCount")
		os.Exit(3)
	}

	for i := 0; i < destroyCount; i++ {
		if destroyFieldsX[i] >= n || destroyFieldsY[i] >= n || destroyFieldsX[i] < 0 || destroyFieldsY[i] < 0 {
			fmt.Println("[Error]: Bad values of destroy coordinates")
			os.Exit(3)
		}
	}
}

func main() {
	fmt.Println("Enter N:")
	fmt.Scan(&n)

	fmt.Println("Enter start position (x;y): ")
	fmt.Scan(&startX, &startY)

	fmt.Println("Enter end position (x;y): ")
	fmt.Scan(&endX, &endY)

	fmt.Println("How many fields are destroy?")
	fmt.Scan(&destroyCount)

	destroyFieldsX = make([]int, destroyCount)
	destroyFieldsY = make([]int, destroyCount)

	fmt.Println("Enter the coordinates of the cells to be cut: ")
	for i := 0; i < destroyCount; i++ {
		fmt.Scan(&destroyFieldsX[i])
		fmt.Scan(&destroyFieldsY[i])
	}

	CheckingArguments()
	RunBruteForce()
}