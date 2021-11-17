package graph

/*
	Implements basic methods for the graph object to obtain connected components
*/

func (g *Graph) DFS(u Node, iSize *int) {
	for _, v := range g.edges[u] {
		if g.visited[v] == false {
			g.visited[v] = true
			*iSize += 1
			g.DFS(v, iSize)
		}
	}
}

func (g *Graph) ConnectedComponents() map[int]int {

	// hash map - {component_size : count}
	hmSizes := make(map[int]int)

	for n := range g.edges {
		if g.visited[n] == false {
			var iSize int = 0
			g.DFS(n, &iSize)
			hmSizes[iSize] += 1
		}
	}

	return hmSizes
}

func (g *Graph) GetGiantConnectedComponentSize() int {
	var components map[int]int = g.ConnectedComponents()

	var iGCCSize int = 0

	for n := range components {
		if n > iGCCSize {
			iGCCSize = n
		}
	}
	return iGCCSize
}
