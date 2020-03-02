package main

import "container/list"

var (
	dx = [8]int {1, 1, -1, -1, 2, 2, -2, -2}
	dy = [8]int {2, -2, 2, -2, 1, -1, 1, -1}
	l = list.New()
)

type Vertex struct {
	x int
	y int
}

func (v *Vertex) Create(x, y int) {
	v.x = x
	v.y = y
}

func GetFromList() *Vertex {
	e := l.Front()
	v := e.Value.(*Vertex)
	return v
}

func IsValidMove(v *Vertex) bool {
	if v.x < 0 || v.x >= n || v.y < 0 || v.y >= n {
		return false
	}

	return true
}

func BFS(chess [][]int) bool {
	vStart := new(Vertex)
	vStart.Create(startX, startY)
	current := 1
	l.PushBack(vStart)

	chess[startX][startY] = -1
	chess[endX][endY] = -2

	PrintChess(chess)
	for l.Len() != 0 {
		node := GetFromList()
		l.Remove(l.Front())
		for k := 0; k < 8; k++ {
			next := new(Vertex)
			next.Create(node.x + dx[k], node.y + dy[k])

			if !IsValidMove(next) {
				continue
			}

			if chess[next.x][next.y] != -3 && 
				chess[next.x][next.y] == 0 {
					if chess[next.x][next.y] == -2 {
						return true
					}
					chess[next.x][next.y] = current
					current++
					l.PushBack(next)
				}
		}
	}
	return false
}