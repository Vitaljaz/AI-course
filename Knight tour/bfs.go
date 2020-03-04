package main

import "fmt"

func RestorePathBFS(vertex *Vertex, chess [][]int) (path [][]int, pathCount int) {
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

	path = chess
	pathCount = i
	return
}

func BFS(chess [][]int) (ok bool, finish *Vertex, count int) {
	ClearList()

	ok = false
	chess[startX][startY] = -1

	vStart := new(Vertex)
	vStart.Create(startX, startY)
	l.PushBack(vStart)

	current := 0

	for l.Len() != 0 {
		node := GetFromListFront()
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
					
					chess[next.x][next.y] = current
					current++

					if next.x == endX && next.y == endY {
						ok = true
						finish = next
						count = current
						return
					}

					l.PushBack(next)
				}
		}
	}
	 return
}