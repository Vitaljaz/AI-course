package main

func BruteForce(chess [][]int, current int, x int, y int) bool {
	if x >= n || y >= n || x < 0 || y < 0 {
		return false
	}

	if x == endX && y == endY {
		chess[x][y] = -2
		return true
	}
	
	if chess[x][y] == -3 {
		return false
	}

	if chess[x][y] > 0 {
		return false
	}

	current++
	chess[x][y] = current

	if current == n * n {
		return true
	}
	if BruteForce(chess, current, x + 2, y + 1) {
		return true
	}
	if BruteForce(chess, current, x + 2, y - 1) {
		return true
	}
	if BruteForce(chess, current, x - 2, y + 1) {
		return true
	}
	if BruteForce(chess, current, x - 2, y - 1) {
		return true
	}
	if BruteForce(chess, current, x + 1, y + 2) {
		return true
	}
	if BruteForce(chess, current, x + 1, y - 2) {
		return true
	}
	if BruteForce(chess, current, x - 1, y + 2) {
		return true
	}
	if BruteForce(chess, current, x - 1, y - 2) {
		return true
	}

	chess[x][y] = 0
	current--
	return false
}