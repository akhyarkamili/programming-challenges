package main

import (
	"slices"
)

type point struct {
	x int
	y int
}

type queue struct {
	points []point
}

func (q *queue) enqueue(p point) {
	q.points = append(q.points, p)
}

func (q *queue) dequeue() point {
	p := q.points[0]
	q.points = slices.Delete(q.points, 0, 1)
	return p
}

func (q *queue) len() int {
	return len(q.points)
}

func (q *queue) peek(position int) point {
	return q.points[position]
}

func numIslands(grid [][]byte) int {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if string(grid[y][x]) == "1" {
				nuke(grid, point{x, y})
			}
		}
	}

	islands := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if string(grid[y][x]) == "1" {
				islands++
			}
		}
	}

	return islands
}

func nuke(grid [][]byte, starter point) {
	var visited = make(map[point]bool)
	var toVisit queue

	toVisit.enqueue(starter)
	for toVisit.len() > 0 {
		var possibleNeighbors []point
		current := toVisit.dequeue()
		if _, ok := visited[current]; ok {
			continue
		}
		visited[current] = true

		x, y := current.x, current.y
		possibleNeighbors = []point{{x + 1, y}, {x, y + 1}, {x - 1, y}, {x, y - 1}}
		for _, n := range possibleNeighbors {
			if isValidUnvisited(visited, n, len(grid), len(grid[0])) &&
				string(grid[n.y][n.x]) == "1" {
				erase(grid, n)
				toVisit.enqueue(n)
			}
		}
	}
}

func erase(grid [][]byte, n point) {
	var eraser = []byte("0")
	grid[n.y][n.x] = eraser[0]
}

func isLand(grid [][]byte, p point) bool {
	return string(grid[p.y][p.x]) == "1"
}

func isValidUnvisited(visited map[point]bool, n point, yLimit int, xLimit int) bool {
	_, ok := visited[n]
	return n.x >= 0 && n.x < xLimit && n.y >= 0 && n.y < yLimit && !ok
}
