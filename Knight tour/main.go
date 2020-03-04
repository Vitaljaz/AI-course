package main

import (
	"fmt"
	"time"
	"os"
	"container/list"
)

var (
	n 	             int
	startX, startY   int
	endX, endY       int
	destroyCount     int
	destroyFieldsX []int
	destroyFieldsY []int
		
	dx = [8]int {1, 1, -1, -1, 2, 2, -2, -2}
	dy = [8]int {2, -2, 2, -2, 1, -1, 1, -1}

	l = list.New()
)

func RunBruteForce() {
	fmt.Println("\n[BruteForce]:")
	chess := CreateChess()
	timer := time.Now()
	
	if BruteForce(chess, 0, startX, startY) {
		fmt.Printf("[BruteForce]: Elapsed time: %v\n", time.Since(timer))
		PrintChess(chess)
	} else {
		fmt.Println("[BruteForce]: Impossible!")
	}
}

func RunBFS() {
	fmt.Println("\n[BFS]:")
	chess := CreateChess()
	timer := time.Now()

	ok, vertex, count := BFS(chess)

	if ok {
		fmt.Printf("[BFS]: Elapsed time: %v\n", time.Since(timer))
		path, pathCount := RestorePathBFS(vertex, CreateChess())
		PrintChess(path)
		fmt.Println("[BFS]: P =", float64(pathCount) / float64(count))
	} else {
		fmt.Println("[BFS]: Impossible!")
		PrintChess(chess)
	}
}

func RunDFS() {
	fmt.Println("\n[DFS]:")
	chess := CreateChess()
	timer := time.Now()

	ok, vertex, count := DFS(chess)

	if ok {
		fmt.Printf("[DFS]: Elapsed time: %v\n", time.Since(timer))
		path, pathCount := RestorePathBFS(vertex, CreateChess())
		PrintChess(path)
		fmt.Println("[DFS]: P =", float64(pathCount) / float64(count))
	} else {
		fmt.Println("[DFS]: Impossible!")
		PrintChess(chess)
	}
}

func ClearList() {
	for l.Len() != 0 {
		l.Remove(l.Back())
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

func CreateChess() [][]int {
	chess := make([][]int, n)
	for i := range chess {
		chess[i] = make([]int, n)
	}

	if destroyCount > 0 {
		for i := 0; i < destroyCount; i++ {
			chess[destroyFieldsX[i]][destroyFieldsY[i]] = -3
		}
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

	if destroyCount > 0 {
		fmt.Println("Enter the coordinates of the cells to be cut: ")
		for i := 0; i < destroyCount; i++ {
			fmt.Scan(&destroyFieldsX[i])
			fmt.Scan(&destroyFieldsY[i])
		}
	}

	CheckingArguments()
	RunBruteForce()
	RunBFS()
	RunDFS()
}