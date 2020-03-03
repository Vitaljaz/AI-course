package main

var (
	dx = [8]int {1, 1, -1, -1, 2, 2, -2, -2}
	dy = [8]int {2, -2, 2, -2, 1, -1, 1, -1}
)

type Vertex struct {
	x int
	y int
}

func (v *Vertex) Create(x, y int) {
	v.x = x
	v.y = y
}

func IsValidMove(v *Vertex) bool {
	if v.x < 0 || v.x >= n || v.y < 0 || v.y >= n {
		return false
	}
	return true
}