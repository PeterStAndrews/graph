package graph

import (
	"math/rand"
	"time"
)

/* takes a degree sequence and creates a graph object according to the configuration model*/
func ConfigurationModel(ns *[]int) *Graph {

	g := Graph{}

	// append nodes to configuration model array
	var config_array []Node
	var nodes []Node
	for i, c := range *ns {
		var n Node = Node(i)
		nodes = append(nodes, n)
		for j := 0; j < c; j++ {
			config_array = append(config_array, n)
		}
	}

	g.AddNodesFrom(nodes)

	// Fisher–Yates shuffle
	rand.Seed(time.Now().UnixNano())
	for i := len(config_array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		config_array[i], config_array[j] = config_array[j], config_array[i]
	}

	// add edges to graph
	for i := 0; i < len(config_array); i += 2 {

		var n1 Node = config_array[i]
		var n2 Node = config_array[i+1]

		if n1 == n2 {
			continue
		}

		g.AddEdge(n1, n2)
	}

	return &g
}
