package graph

import (
	"math/rand"
	"time"
)

/* Implements bond percolation method over graph */

func linspace(start, stop float64, N int) []float64 {
	step := (stop - start) / float64(N)
	rnge := make([]float64, N, N)
	i := 0
	for x := start; x < stop; x += step {
		rnge[i] = x
		i += 1
	}
	return rnge
}

// performs bond percolation on graph, returns size of GCC
func BondPercolation(g *Graph, T float64) float64 {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var es []*Edge
	for k, v := range g.edges {
		for _, n := range v {
			if k < n { // this ensures we don't add each edges twice
				if T < r.Float64() {
					e := Edge{k, n}
					es = append(es, &e)
				}
			}
		}
	}

	g.RemoveEdgesFrom(&es)
	return float64(g.GetGiantConnectedComponentSize()) / float64(len(g.edges))
}

func BondPercolationExperiment() *[]float64 {

	var Ts []float64 = linspace(0.0, 1.0, 20)
	var Ss []float64
	for _, T := range Ts {

		var p Poisson    // poisson degree distribution
		p.kmean = 2.5    // average degree
		const N = 100000 // number of nodes

		ns := p.generateSamples(N)
		Ss = append(Ss, BondPercolation(ConfigurationModel(&ns), T))
	}

	return &Ss
}
