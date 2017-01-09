package dna

import (
	"math/rand"
	"time"
	"fmt"
)

const MUTATION_RATE = 0.01

// The structure for DNA, which is a list of genes
type DNA struct {
	genes        []float64
}

// Create a DNA sequence with a given list of genes
func CreateDNA(genes []float64) *DNA {
	g := make([]float64, len(genes))
	for i := 0; i < len(genes); i++ {
		if genes[i] < 0 {
			g[i] = 0.
		} else if genes[i] > 1 {
			g[i] = 1.
		}	else {
			g[i] = genes[i]
		}
	}

	return &DNA{g}
}

// Create a DNA seqence of length n, chosen randomly
func CreateRandDNA(n int) *DNA {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	g := make([]float64, n)
	for i := 0; i < n; i++ {
		g[i] = r.Float64()
	}
	return &DNA{g}
}

// Create a crossover strand
func (self *DNA) Crossover(partner *DNA) *DNA {
	l := self.Length()
	g := make([]float64, l)
	partnerLength := partner.Length()
	mid := l / 2
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	rand := r.Float64()

	for i := 0; i < l; i++ {
		if rand < 0.5 {
			if i < mid && i < partnerLength {
				g[i] = partner.genes[i]
			} else {
				g[i] = self.genes[i]
			}
		} else {
			if i < mid || i >= partnerLength {
				g[i] = self.genes[i]
			} else {
				g[i] = partner.genes[i]
			}
		}
	}

	return CreateDNA(g)
}

// Mutate a strand
func (self *DNA) Mutate() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < self.Length(); i++ {
		if r.Float64() < MUTATION_RATE {
			this.genes[i] = r.Float64
		}
	}
}

// TODO calculate the fitness of the DNA
func (self *DNA) Fitness() float64 {
	return 0.
}

// Get number of genes
func (self *DNA) Length() int {
	return len(self.genes)
}

// Print DNA properly
func (self *DNA) String() string {
	s := "["
	for i := 0; i < self.Length(); i++ {
		if i != 0 {
			s += ", "
		}
		// %.2f => float number with precision 2
		s += fmt.Sprintf("%.2f", self.genes[i])
	}
	s += "]"
	return s
}
