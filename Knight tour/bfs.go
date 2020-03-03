package main

import (
	"container/list"
	"fmt"
)

var (
	l = list.New()
)

func GetFromList() *Vertex {
	e := l.Front()
	v := e.Value.(*Vertex)
	return v
}

func RestorePath(vertex *Vertex, chess [][]int) [][]int {
	chess[endX][endY] = -2
	v := vertex.parent
	i := 1
	fmt.Println("[BFS]: Reverse Path: ")
	fmt.Printf("(%d;%d) ", endX, endY)
	for v != nil {
		chess[v.x][v.y] = i
		i++
		fmt.Printf("(%d;%d) ", v.x, v.y)
		v = v.parent
	}
	fmt.Printf("\n")
	return chess
}

func BFS(chess [][]int) (can bool, finish *Vertex) {
	vStart := new(Vertex)
	vStart.Create(startX, startY)
	l.PushBack(vStart)

	can = false
	chess[startX][startY] = -1

	for l.Len() != 0 {
		node := GetFromList()
		l.Remove(l.Front())

		for k := 0; k < 8; k++ {
			next := new(Vertex)
			next.Create(node.x + dx[k], node.y + dy[k])

			if !IsValidMove(next) {
				continue
			}

			if chess[next.x][next.y] == 0 &&
				chess[next.x][next.y] != -3 {
					next.parent = node

					if next.x == endX && next.y == endY {
						can = true
						finish = next
						return
					}

					l.PushBack(next)
				}
		}
	}
	 return
}