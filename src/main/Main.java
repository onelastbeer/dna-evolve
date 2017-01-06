package main;

import java.util.ArrayList;
import java.util.List;

import dna.DNA;

public class Main {
	
	private static final int POPULATION_SIZE = 100;
	private static final int LOOPS = 10000;
	private static final int GENES_NUMBER = 5;
	private static List<DNA> population = new ArrayList<DNA>();
	private static List<DNA> DNAPool = new ArrayList<DNA>();

	public static void main(String[] args) {
		
		long startTime = System.currentTimeMillis();
		
		init();
		//printPopulation();
		
		for(int i = 0; i < LOOPS; ++i) {
			nextDNAPool();
			nextGen();
			//printPopulation();
			//printPopulationFitness();
		}
		
		printPopulationFitness();
		
		long endTime = System.currentTimeMillis();

		System.out.println("Total run time : " + (endTime-startTime)/1000.0 + " s");
		System.out.println("Time per generation : " + (endTime-startTime)/LOOPS/1000.0 + " s");
		
	}

	private static void init() {
		for (int i = 0; i < POPULATION_SIZE; ++i) {
			population.add(new DNA(GENES_NUMBER));
		}
	}
	
	private static void mutate() {
		for(DNA dna : population) {
			dna.mutate();
		}
	}
	
	private static void nextDNAPool() {
		DNAPool = new ArrayList<DNA>();
		for(DNA dna : population) {
			int fitness = (int)(dna.fitness()*100);
			for(int i = 0; i < fitness; ++i) {
				DNAPool.add(dna);
			}
		}
	}
	
	private static void nextGen() {
		population = new ArrayList<DNA>(); 
		for (int i = 0; i < POPULATION_SIZE; ++i) {
			DNA parent1 = DNAPool.get((int)(Math.random()*DNAPool.size()));
			DNA parent2 = DNAPool.get((int)(Math.random()*DNAPool.size()));
			population.add(parent1.crossover(parent2));
		}
		mutate();
	}
	
	private static void printPopulation() {
		for(DNA dna : population) {
			System.out.println(dna.toString());
		}
	}
	
	private static void printPopulationFitness() {
		
		double sum = 0;
		DNA best = population.get(0);
		double bestFit = population.get(0).fitness();
		
		for(DNA dna : population) {
			double fit = dna.fitness();
			sum += fit;
			if(fit > bestFit) {
				best = dna;
				bestFit = fit;
			}
		}
		
		System.out.println("Population average fitness : " + sum/population.size());
		System.out.println("Best DNA : " + best.toString() + "- Fitness : " + bestFit);
	}

}
