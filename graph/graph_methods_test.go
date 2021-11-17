package graph

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestComponents(t *testing.T) {
	g := Graph{}

	var n1 Node = 0
	var n2 Node = 1
	var n3 Node = 2
	var n4 Node = 3
	var n5 Node = 4
	var n6 Node = 5
	var n7 Node = 6
	var n8 Node = 7
	var n9 Node = 8

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)
	g.AddNode(n7)
	g.AddNode(n8)
	g.AddNode(n9)

	g.AddEdge(n1, n2)

	g.AddEdge(n3, n4)
	g.AddEdge(n4, n5)
	g.AddEdge(n5, n3)

	g.AddEdge(n6, n7)
	g.AddEdge(n7, n8)
	g.AddEdge(n8, n9)
	g.AddEdge(n9, n6)
	g.AddEdge(n6, n8)
	g.AddEdge(n7, n9)

	var components map[int]int = g.ConnectedComponents()

	if components[2] != 1 || components[3] != 1 || components[4] != 1 {
		t.Error("Error: TestComponents - incorrect component numbers")
	}
}

func TestGiantConnectedComponentSize(t *testing.T) {
	g := Graph{}

	var n1 Node = 0
	var n2 Node = 1
	var n3 Node = 2
	var n4 Node = 3
	var n5 Node = 4
	var n6 Node = 5
	var n7 Node = 6
	var n8 Node = 7
	var n9 Node = 8

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)
	g.AddNode(n7)
	g.AddNode(n8)
	g.AddNode(n9)

	g.AddEdge(n1, n2)

	g.AddEdge(n3, n4)
	g.AddEdge(n4, n5)
	g.AddEdge(n5, n3)

	g.AddEdge(n6, n7)
	g.AddEdge(n7, n8)
	g.AddEdge(n8, n9)
	g.AddEdge(n9, n6)
	g.AddEdge(n6, n8)
	g.AddEdge(n7, n9)

	if g.GetGiantConnectedComponentSize() != 4 {
		t.Error("Error: GCC")
	}
}

func TestComponentsAtScale(t *testing.T) {

	/* functionality test to ensure codebase scales to large graph*/
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	g := Graph{}

	var ns []Node
	var N int = 100000 // large number of nodes
	var E int = 250000 // large number of edges

	for i := 0; i < N; i++ {
		var n Node = Node(i)
		ns = append(ns, n)
	}

	g.AddNodesFrom(ns)

	var es []*Edge
	for i := 0; i < E; i++ {
		var condition bool = false
		for ok := true; ok; ok = !condition {

			var n1 Node = Node(r1.Intn(E))
			var n2 Node = Node(r1.Intn(E))

			if n1 != n2 {
				condition = true
				e := Edge{n1, n2}
				es = append(es, &e)
			}
		}
	}

	g.AddEdgesFrom(&es)

	var components map[int]int = g.ConnectedComponents()

	fmt.Printf("components: %v\n", components)
}
