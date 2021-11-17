package graph

import (
	"math/rand"
	"testing"
)

func TestAddNode(t *testing.T) {

	g := Graph{}

	var n1 Node = 0
	var n2 Node = 1
	var n3 Node = 2

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)

	var want int = 3
	var got int = len(g.edges)

	if got != want {
		t.Errorf("Error: AddNode - got %q, wanted %q", got, want)
	}
}
func TestAddEdge(t *testing.T) {

	g := Graph{}

	var n1 Node = 0
	var n2 Node = 1
	var n3 Node = 2

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)

	g.AddEdge(n1, n2)
	g.AddEdge(n1, n3)

	var want int = 2
	var got int = g.EnumerateEdges()

	if got != want {
		t.Errorf("Error: AddEdge - got %q, wanted %q", got, want)
	}
}
func TestHasEdge(t *testing.T) {

	g := Graph{}

	var n1 Node = 0
	var n2 Node = 1
	var n3 Node = 2

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)

	g.AddEdge(n1, n2)
	g.AddEdge(n1, n3)

	if g.HasEdge(n1, n2) != g.HasEdge(n2, n1) {
		t.Error("Error: HasEdge - 1")
	}
	if !g.HasEdge(n1, n2) || !g.HasEdge(n1, n3) {
		t.Error("Error: HasEdge - 2")
	}
	if g.HasEdge(n2, n3) {
		t.Error("Error: HasEdge - 3")
	}
}
func TestRemoveEdge(t *testing.T) {

	g := Graph{}

	var n1 Node = 0
	var n2 Node = 1
	var n3 Node = 2

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)

	g.AddEdge(n1, n2)
	g.AddEdge(n1, n3)
	g.AddEdge(n2, n3)

	var want int = 3
	var got int = g.EnumerateEdges()

	if got != want {
		t.Errorf("Error: RemoveEdge - got %q, wanted %q", got, want)
	}

	g.RemoveEdge(n1, n2)

	want = 2
	got = g.EnumerateEdges()

	if got != want {
		t.Errorf("Error: RemoveEdge - got %q, wanted %q", got, want)
	}

	if g.HasEdge(n1, n2) {
		t.Error("Error: RemoveEdge - Edge not removed correctly")
	}
}
func TestRemoveEdgesFrom(t *testing.T) {

	g := Graph{}

	var n1 Node = 0
	var n2 Node = 1
	var n3 Node = 2

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)

	g.AddEdge(n1, n2)
	g.AddEdge(n1, n3)
	g.AddEdge(n2, n3)

	e01 := Edge{n1, n2}
	e02 := Edge{n1, n3}
	es := []*Edge{&e01, &e02}

	g.RemoveEdgesFrom(&es)

	if g.HasEdge(e01.u, e01.v) || g.HasEdge(e02.u, e02.v) {
		t.Error("Error: RemoveEdgesFrom - edges weren't removed")
	}

	var want int = 1
	var got int = g.EnumerateEdges()

	if got != want {
		t.Errorf("Error: RemoveEdge - got %q, wanted %q", got, want)
	}
}
func TestAddNodesFrom(t *testing.T) {
	g := Graph{}

	var ns []Node
	var N int = 100

	for i := 0; i < N; i++ {
		var n Node = Node(i)
		ns = append(ns, n)
	}

	g.AddNodesFrom(ns)

	if g.EnumerateNodes() != int(N) {
		t.Error("Error: AddNodesFrom")
	}
}
func TestAddEdgesFrom(t *testing.T) {

	g := Graph{}

	var ns []Node
	var N int = 1000 // number of nodes
	var E int = 1000 // number of edges

	for i := 0; i < N; i++ {
		var n Node = Node(i)
		ns = append(ns, n)
	}

	g.AddNodesFrom(ns)

	var es []*Edge
	for i := 0; i < E; i++ {
		var condition bool = false
		for ok := true; ok; ok = !condition {

			var n1 Node = Node(rand.Intn(E))
			var n2 Node = Node(rand.Intn(E))

			if n1 != n2 {
				condition = true
				e := Edge{n1, n2}
				es = append(es, &e)
			}
		}
	}

	g.AddEdgesFrom(&es)
	if g.EnumerateEdges() != len(es) { // fails if e_ij and e_ji are both in es
		t.Error("Error: AddEdgesFrom")
	}
}
func TestEdgesContainerFromAdjacencyList(t *testing.T) {

	g := Graph{}

	var n1 Node = 0
	var n2 Node = 1
	var n3 Node = 2
	var n4 Node = 3
	var n5 Node = 4

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)

	g.AddEdge(n1, n2)
	g.AddEdge(n2, n3)
	g.AddEdge(n3, n4)
	g.AddEdge(n4, n5)
	g.AddEdge(n1, n5)

	var ns = g.edges[n1]
	es := g.EdgesContainerFromAdjacencyList(n1, ns)

	if len(es) != 2 {
		t.Error("Error: EdgesContainerFromAdjacencyList()")
	}
	if es[0].u != 0 || es[0].v != 1 {
		t.Error("Error: EdgesContainerFromAdjacencyList()")
	}
	if es[1].u != 0 || es[1].v != 4 {
		t.Error("Error: EdgesContainerFromAdjacencyList()")
	}
}
func TestHasNode(t *testing.T) {
	g := Graph{}

	var ns []Node
	var N int = 100

	for i := 0; i < N; i++ {
		var n Node = Node(i)
		ns = append(ns, n)
	}

	g.AddNodesFrom(ns)

	if !g.HasNode(ns[rand.Intn(N)]) {
		t.Error("Error: HasNode")
	}
}
func TestRemoveNode(t *testing.T) {
	g := Graph{}

	var n1 Node = 0
	var n2 Node = 1
	var n3 Node = 2
	var n4 Node = 3
	var n5 Node = 4

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)

	g.AddEdge(n1, n2)
	g.AddEdge(n2, n3)
	g.AddEdge(n3, n4)
	g.AddEdge(n4, n5)
	g.AddEdge(n1, n5)

	g.RemoveNode(n1)

	if g.HasNode(n1) {
		t.Error("Error: RemoveNode - node still present")
	}
	if g.HasEdge(n1, n2) {
		t.Error("Error: RemoveNode - still has edge")
	}
	if g.EnumerateNodes() != 4 {
		t.Error("Error: RemoveNode - node count incorrect")
	}
	if g.EnumerateEdges() != 3 {
		t.Error("Error: RemoveNode - edge count incorrect")
	}
}
