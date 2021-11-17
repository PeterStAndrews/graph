package graph

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	var p Poisson
	if p.factorial(9) != 362880 {
		t.Error("Error: factorial")
	}
}

func TestRandInt(t *testing.T) {

	const min int = 10
	const max int = 50

	var p Poisson
	for i := 0; i < 1000; i++ {
		var r = p.randInt(min, max)
		if r < min || r > max {
			t.Errorf("Error: randInt %d, %d", min, max)
		}
	}
}

func TestGenerateSamples(t *testing.T) {

	var p Poisson
	const N = 100000

	p.kmean = 3.0
	ns := p.generateSamples(N)

	var sum int = 0
	for n := range ns {
		sum += n
	}

	if sum%2 != 0 {
		t.Error("Error: generateSamples")
	}
}
