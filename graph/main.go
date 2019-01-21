package main

import (
	"fmt"
)

type Graph struct {
	AdjacencyList map[string][]string
}

/**
 * 加入節點。
 */
func (g *Graph) AddVertex(vertex string) {
	if _, ok := g.AdjacencyList[vertex]; ok {
		return
	}

	g.AdjacencyList[vertex] = make([]string, 0)
}

/**
 * 加入邊。
 */
func (g *Graph) AddEdge(vertex1, vertex2 string) {
	g.AdjacencyList[vertex1] = append(g.AdjacencyList[vertex1], vertex2)
	g.AdjacencyList[vertex2] = append(g.AdjacencyList[vertex2], vertex1)
}

/**
 * 移除邊。
 */
func (g *Graph) RemoveEdge(vertex1, vertex2 string) {
	g.AdjacencyList[vertex1] = g.remove(g.AdjacencyList[vertex1], vertex2)
	g.AdjacencyList[vertex2] = g.remove(g.AdjacencyList[vertex2], vertex1)
}

/**
 * 移除節點。
 */
func (g *Graph) RemoveVertex(vertex string) {
	for len(g.AdjacencyList[vertex]) > 0 {
		g.RemoveEdge(vertex, g.AdjacencyList[vertex][0])
	}

	delete(g.AdjacencyList, vertex)
}

/**
 * 深度優先搜尋。
 */
func (g *Graph) DepthFirstSearch(start string) []string {
	result := make([]string, 0)
	visited := make(map[string]bool)

	var dfs func(vertex string)
	dfs = func(vertex string) {
		if _, ok := visited[vertex]; ok {
			return
		}

		visited[vertex] = true
		result = append(result, vertex)

		for _, v := range g.AdjacencyList[vertex] {
			dfs(v)
		}
	}

	dfs(start)
	return result
}

func (g *Graph) remove(vertexs []string, vertex string) []string {
	index := -1

	for i, v := range vertexs {
		if v == vertex {
			index = i
		}
	}

	if index != -1 {
		vertexs = append(vertexs[:index], vertexs[index+1:]...)
	}

	return vertexs
}

func main() {
	//     A
	//   /  \
	//  B    C
	//  |    |
	//  D----E
	//  \   /
	//    F
	graph := Graph{AdjacencyList: make(map[string][]string)}
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddVertex("F")

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "E")
	graph.AddEdge("D", "E")
	graph.AddEdge("D", "F")
	graph.AddEdge("E", "F")

	fmt.Println(graph.DepthFirstSearch("A"))
}
