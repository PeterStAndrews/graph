package graph

import (
	"math"
	"math/rand"
	"time"
)

/*
	Implements methods to create a joint degree sequence for the configuration model
*/

type distribution interface {
	factorial(n int) uint64
	rand(...interface{}) float64 // variadic arguments
	generateSamples(N int) []float64
	randInt(min int, max int) int
}

type Poisson struct {
	kmean float64
	r     *rand.Rand
}

func (p *Poisson) factorial(n int) uint64 {
	var factVal uint64 = 1
	if n < 0 {
		szError := "Error: factorial function of negative number"
		panic(szError)
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}
	}
	return factVal
}

func (p *Poisson) rand(k float64) float64 {
	return math.Pow(p.kmean, k) * (math.Exp(-p.kmean) / float64(p.factorial(int(k))))
}

func (p *Poisson) randInt(min int, max int) int {

	if p.r == nil {

		s1 := rand.NewSource(time.Now().UnixNano())
		p.r = rand.New(s1)
	}
	return min + p.r.Intn(max-min)
}

func (p *Poisson) generateSamples(N int) []int {

	s1 := rand.NewSource(time.Now().UnixNano())
	p.r = rand.New(s1)

	const min int = 0
	var max int = 4 * int(p.kmean) // is this window wide enough?

	if max < 1 { // bug, for small mean degree, the maximum window size is zero
		max = 1 // this causes p.randInt to be (0,0) which panics.
	}

	var t int = 0
	var ns []int
	for i := 0; i < N; i++ {
		for {
			var k float64 = float64(p.randInt(min, max))

			if p.r.Float64() < p.rand(k) {
				ns = append(ns, int(k))
				t += int(k)
				break
			}
		}
	}

	for {

		if t%2 == 0 {
			break
		}

		var i int = p.randInt(0, len(ns)-1)
		ns[i] = ns[len(ns)-1]
		ns = ns[:len(ns)-1]

		for {
			var k float64 = float64(p.randInt(min, max))
			if p.r.Float64() < p.rand(k) {
				ns = append(ns, int(k))
				t += int(k)
				break
			}
		}
	}
	return ns
}
