package main

import "fmt"

type set map[int]bool

type graph struct {
	adj map[int]set
}

func create() *graph {
	var g = graph{}
	g.adj = make(map[int]set)
	return &g
}

func (g *graph) add(node1 int, node2 int) {
	if g.adj[node1] == nil {
		g.adj[node1] = make(set)
	}
	g.adj[node1][node2] = true

	if g.adj[node2] == nil {
		g.adj[node2] = make(set)
	}
	g.adj[node2][node1] = true
}

func dfs(g *graph, current int, visited set, visitFunction func(int)) {
	if _, seen := visited[current]; seen {
		return
	}

	visited[current] = true
	visitFunction(current)

	for neighbour := range g.adj[current] {
		dfs(g, neighbour, visited, visitFunction)
	}
}

func main() {
	graph := create()
	graph.add(1, 2)
	graph.add(2, 3)
	graph.add(2, 4)

	visited := make(set)

	dfs(graph, 1, visited, func(node int) {
		fmt.Print(node, " ")
	})
}
