package graph

import (
	"fmt"
	"testing"
)

func TestBondPercolation(t *testing.T) {

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
	g.AddEdge(n1, n2)
	g.AddEdge(n3, n4)
	g.AddEdge(n4, n5)
	g.AddEdge(n5, n3)
	g.AddEdge(n5, n6)
	g.AddEdge(n6, n7)
	g.AddEdge(n7, n8)
	g.AddEdge(n8, n9)
	g.AddEdge(n9, n6)
	g.AddEdge(n6, n8)
	g.AddEdge(n7, n9)

	fmt.Printf("Size of GCC = %f \n", BondPercolation(&g, 0.5))
}

func TestBondPercolationAtScaleLowT(t *testing.T) {

	var p Poisson    // poisson degree distribution
	p.kmean = 3      // average degree
	const N = 100000 // number of ndoes

	ns := p.generateSamples(N)

	const T float64 = 1e-3
	S := BondPercolation(ConfigurationModel(&ns), T)
	fmt.Printf("Size of GCC = %f \n", S)

	if S > 0.01 {
		t.Error("Error - bond percolation low T")
	}
}

func TestBondPercolationAtScaleHighT(t *testing.T) {

	var p Poisson    // poisson degree distribution
	p.kmean = 3      // average degree
	const N = 100000 // number of ndoes

	ns := p.generateSamples(N)

	const T float64 = 0.999
	S := BondPercolation(ConfigurationModel(&ns), T)
	fmt.Printf("Size of GCC = %f \n", S)

	if S < 0.9 {
		t.Error("Error - bond percolation hight T")
	}

}
