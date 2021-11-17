package graph

/*
	Implements basic structure for a graph object
*/

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

// Node
type Node int

// Edge
type Edge struct {
	u, v Node
}

// containers for nodes and edges
type Nodes []*Node
type Edges map[Node][]Node

// abstract interface for graph class
type GraphInterface interface {
	AddNode(u *Node) // AddNode adds a node to the graph and sets unvisited

	AddEdge(u, v *Node) // AddEdge adds an edge to the graph

	AddNodesFrom(ns []*Node) // Add nodes from container of nodes

	AddEdgesFrom(es []*Edge) // Add edges from container of edges

	RemoveEdge(u, v *Node) // Removes edge from graph

	RemoveEdgesFrom(es *[]*Edges) // Remove edges from a container of edges

	EdgesContainerFromAdjacencyList(n *Node, ns *[]*Node) []*Edge // creates type(Edges) container from adjacency list

	RemoveNode(u *Node) // remove node from graph

	RemoveNodeFrom(ns []*Node) // remove nodes from container of nodes

	HasEdge(u, v *Node) bool // check if edge is in the graph

	HasNode(u Node) bool // Check if a node is in the graph

	EnumerateNodes() int // return a count of the number of nodes

	EnumerateEdges() int // return a count of the number of edges

	DFS(u *Node, iSize *int) // recursive depth-first search routine

	ConnectedComponents() map[int]int // return hash-map of connected components {component_size : count}

	GetGiantConnectedComponentSize() int // return size of largest connected component
}

// Graph structure
type Graph struct {
	edges   Edges
	visited map[Node]bool
}

func (g *Graph) AddNode(u Node) {

	if g.edges == nil {
		g.edges = make(Edges)
	}

	if g.HasNode(u) {
		return
	}

	g.edges[u] = []Node{}

	if g.visited == nil {
		g.visited = make(map[Node]bool)
	}

	g.visited[u] = false
}

func (g *Graph) AddEdge(u, v Node) {

	if g.HasEdge(u, v) {
		return
	}

	if g.edges == nil {
		g.edges = make(Edges)
	}

	g.edges[u] = append(g.edges[u], v)
	g.edges[v] = append(g.edges[v], u)
}

func (g *Graph) RemoveEdge(u, v Node) {

	if g.HasEdge(u, v) != true {
		return
	}

	for i, j := range g.edges[u] {
		if j == v {
			g.edges[u][i] = g.edges[u][len(g.edges[u])-1]
			g.edges[u] = g.edges[u][:len(g.edges[u])-1]
		}
	}

	for i, j := range g.edges[v] {
		if j == u {
			g.edges[v][i] = g.edges[v][len(g.edges[v])-1]
			g.edges[v] = g.edges[v][:len(g.edges[v])-1]
		}
	}
}

func (g *Graph) AddNodesFrom(ns []Node) {
	for _, n := range ns {
		g.AddNode(n)
	}
}

func (g *Graph) AddEdgesFrom(es *[]*Edge) {
	for _, e := range *es {
		g.AddEdge(e.u, e.v)
	}
}

func (g *Graph) RemoveEdgesFrom(es *[]*Edge) {
	for _, e := range *es {
		g.RemoveEdge(e.u, e.v)
	}
}

func (g *Graph) HasNode(u Node) bool {

	if _, ok := g.edges[u]; ok {
		return true
	}
	return false
}

func (g *Graph) HasEdge(u, v Node) bool {
	var retU = false
	var retV = false

	for _, j := range g.edges[u] {
		if j == v {
			retU = true
		}
	}

	for _, j := range g.edges[v] {
		if j == u {
			retV = true
		}
	}

	if retU != retV {
		var szErr string = "Error in edge handling"
		panic(szErr)
	}
	return retU && retV
}

func (g *Graph) EnumerateEdges() int {
	var nEdges int = 0
	for _, es := range g.edges {
		nEdges += len(es)
	}
	return nEdges / 2
}

func (g *Graph) EnumerateNodes() int {
	return len(g.edges)
}

func (g *Graph) EdgesContainerFromAdjacencyList(n Node, ns []Node) []*Edge {
	var es []*Edge

	for _, k := range ns {
		e := Edge{n, k}
		es = append(es, &e)
	}

	return es
}

func (g *Graph) RemoveNode(u Node) {
	es := g.EdgesContainerFromAdjacencyList(u, g.edges[u])
	g.RemoveEdgesFrom(&es)
	delete(g.edges, u)
}
