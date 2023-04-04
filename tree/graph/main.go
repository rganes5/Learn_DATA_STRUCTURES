package main

import (
	"fmt"
)

func main() {
	list := &Graph{}
	for i := 1; i < 8; i++ {
		list.addVertex(i)
	}
	list.addEdge(1, 3)
	list.addEdge(1, 2)
	list.addEdge(3, 1)
	list.addEdge(7, 5)
	list.addEdge(5, 6)
	list.addEdge(6, 7)
	list.addEdge(1, 7)
	list.Print()
	fmt.Printf("\nBFS ")
	list.bfsTravers(1)
	fmt.Println("DFS", list.dfsTravers(1))

}

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

type Queue []int

func (g *Graph) addVertex(key int) {
	g.vertices = append(g.vertices, &Vertex{key: key})
}

func (g *Graph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge %v..>%v", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}

}

// DFS TRAVERSAL
func (g *Graph) dfsTravers(start int) []int {
	visited := make(map[int]bool)
	stack := []int{start}
	result := []int{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			visited[node] = true
			result = append(result, node)

			for _, neighbor := range g.getVertex(node).adjacent {
				stack = append(stack, neighbor.key)
			}
		}
	}

	return result
}

// BFS TRAVERSAL
func (g *Graph) bfsTravers(start int) {
	q := Queue{}
	q = append(q, start)
	visited := make(map[int]bool)

	for len(q) != 0 {
		vertex := q[0]
		q = q[1:]
		if !visited[vertex] {
			visited[vertex] = true
			fmt.Printf("%d ", vertex)

			for _, neighbor := range g.vertices {
				if !visited[neighbor.key] {
					q = append(q, neighbor.key)
				}
			}
		}
	}
	fmt.Println()
}
