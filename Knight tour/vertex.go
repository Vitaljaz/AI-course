package main

type Vertex struct {
	x int
	y int

	parent *Vertex
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