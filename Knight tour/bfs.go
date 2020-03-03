package main

import (
	"container/list"
)

var (
	l = list.New()
)

func GetFromList() *Vertex {
	e := l.Front()
	v := e.Value.(*Vertex)
	return v
}

func BFS(chess [][]int) bool {
	vStart := new(Vertex)
	vStart.Create(startX, startY)
	current := 1
	l.PushBack(vStart)

	for l.Len() != 0 {
		node := GetFromList()
		l.Remove(l.Front())

		for k := 0; k < 8; k++ {
			next := new(Vertex)
			next.Create(node.x + dx[k], node.y + dy[k])
		}
	}

	 return false
}