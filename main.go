package main

import (
	"./dna"
	"fmt"
)

func main() {
	fmt.Println(dna.CreateRandDNA(4).Crossover(dna.CreateRandDNA(4)))
	d := dna.CreateRandDNA(4)
	d.Mutate()
	fmt.Println(d)
	fmt.Println(d.Fitness())
}
