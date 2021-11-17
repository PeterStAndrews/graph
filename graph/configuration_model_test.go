package graph

import (
	"testing"
)

func TestConfigurationModel(t *testing.T) {

	var ns = []int{1, 1}

	var g *Graph = ConfigurationModel(&ns)

	if g.EnumerateNodes() != 2 {
		t.Errorf("Error: ConfigurationModel - node count %d", g.EnumerateNodes())
	}

	if g.EnumerateEdges() != 1 {
		t.Errorf("Error: ConfigurationModel - edge count %d", g.EnumerateEdges())
	}
}

func TestConfigurationModelPoisson(t *testing.T) {

	var p Poisson    // poisson degree distribution
	p.kmean = 3.0    // average degree
	const N = 100000 // number of ndoes

	ns := p.generateSamples(N)

	ConfigurationModel(&ns)
}
