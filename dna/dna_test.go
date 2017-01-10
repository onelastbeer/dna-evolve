package dna

import "testing"

var genes1 = []float64{0.3422, 0.33432, 0.42325, 0.8333, 0.5, 0.5}
var genes2 = []float64{0.1324213, 0.6532, 0.454, 0.8425, 0.521, 0.9523}
var shortGenes= []float64{0.234, 0.74321}
var outOfBoundsGenes = []float64{1.3422, 0.33432, 0.42325, -0.8333, -0.5, -0.5}
var dna1 = CreateDNA(genes1)
var dna2 = CreateDNA(genes2)
var shortDna = CreateDNA(shortGenes)
var outOfBoundsDna = CreateDNA(outOfBoundsGenes)
var mutatedDna1 = CreateDNA(genes1)
var mutatedDna2 = CreateDNA(genes1)

/*func genesArray1() [6]float64 = {
  return [6]{0.3422, 0.33432, 0.42325, 0.8333, 0.5, 0.5}
}*/

func TestLength(t *testing.T) {
  if dna1.Length() != 6 {
    t.Error("Length() method returned", dna1.Length(), "instead of", 6)
  }
}

func TestGene(t *testing.T) {
  if dna1.genes[0] != genes1[0] {
    t.Error("genes access returned", dna1.genes[0], "intead of", genes1[0])
  }
}

func TestString(t *testing.T) {
  s1 := "[0.34, 0.33, 0.42, 0.83, 0.50, 0.50]"
  s2 := "[1.00, 0.33, 0.42, 0.00, 0.00, 0.00]"
  if dna1.String() != s1 {
    t.Error("String() method returned", dna1.String(), "instead of", s1)
  }
  if outOfBoundsDna.String() != s2 {
    t.Error("String() method returned", outOfBoundsDna.String(), "instead of", s2)
  }
}

func TestRandDNA(t *testing.T) {
  d := CreateRandDNA(4)
  if d.Length() != 4 {
    t.Error("Length() method on random DNA returned", d.Length(), "instead of", 4)
  }
  for i := 0; i < d.Length(); i++ {
    if d.genes[i] < 0 || d.genes[i] >= 1 {
      t.Error("Invalid value", d.genes[i],"in random DNA")
    }
  }
}

func TestCrossover(t *testing.T) {
  c1 := dna1.Crossover(dna2)
  c2 := dna1.Crossover(shortDna)
  c3 := shortDna.Crossover(dna1)

  if c1.Length() != 6 {
    t.Error("Crossover 1 has Length()", c1.Length(), "instead of", 6)
  }

  if c1.String() != "[0.13, 0.65, 0.45, 0.83, 0.50, 0.50]" && c1.String() != "[0.34, 0.33, 0.42, 0.84, 0.52, 0.95]" {
    t.Error("Crossover 1 has result ", c1.String())
  }

  if c2.Length() != 6 {
    t.Error("Crossover 2 has Length()", c1.Length(), "instead of", 6)
  }

  if c2.String() != "[0.34, 0.33, 0.42, 0.83, 0.50, 0.50]" && c1.String() != "[0.23, 0.74, 0.42, 0.83, 0.50, 0.50]" {
    t.Error("Crossover 2 has result ", c1.String())
  }

  if c3.Length() != 2 {
    t.Error("Crossover 3 has Length()", c1.Length(), "instead of", 2)
  }

  if c3.String() != "[0.23, 0.33]" && c3.String() != "[0.34, 0.74]" {
    t.Error("Crossover 3 has result ", c1.String())
  }

}

func TestMutate(t *testing.T) {
  mutatedDna1.MutationRate(1.)
  mutatedDna1.Mutate()
  if mutatedDna1.String() == dna1.String() {
    t.Error("Mutated DNA 1 (rate=1) is the same as not mutated")
  }
  mutatedDna2.MutationRate(0.)
  mutatedDna2.Mutate()
  if mutatedDna2.String() != dna1.String() {
    t.Error("Mutated DNA 2 (rate=0) is not the same as not mutated")
  }
}
