package dna;

import static org.junit.Assert.*;

import java.util.Arrays;
import java.util.List;

import org.junit.Test;

public class DNATest {
	private static final Double[] GENES_ARRAY1 = { 0.3422, 0.33432, 0.42325, 0.8333, 0.5, 0.5 };
	private static final Double[] GENES_ARRAY2 = { 0.1324213, 0.6532, 0.454, 0.8425, 0.521, 0.9523 };
	private static final Double[] SHORT_GENES_ARRAY = { 0.234, 0.74321};
	private static final Double[] OUT_OF_BOUNDS_GENES_ARRAY = { 1.3422, 0.33432, 0.42325, -0.8333, -0.5, -0.5 };
	private static final List<Double> GENES1 = Arrays.asList(GENES_ARRAY1);
	private static final List<Double> GENES2 = Arrays.asList(GENES_ARRAY2);
	private static final List<Double> SHORT_GENES = Arrays.asList(SHORT_GENES_ARRAY);
	private static final List<Double> OFB_GENES = Arrays.asList(OUT_OF_BOUNDS_GENES_ARRAY);
	private static final DNA DNA1 = new DNA(GENES1);
	private static final DNA DNA2 = new DNA(GENES2);
	private static final DNA SHORT_DNA = new DNA(SHORT_GENES);
	private static final DNA OFB_DNA = new DNA(OFB_GENES);
	private static final DNA MUT_DNA1 = new DNA(GENES1);
	private static final DNA MUT_DNA2 = new DNA(GENES1);

	@Test
	public void lengthTest() {
		assertEquals(DNA1.length(), DNA1.getGenes().size());
	}
	
	@Test
	public void getGeneTest() {
		assertEquals(DNA1.getGene(0), DNA1.getGenes().get(0), 0.01);
	}
	
	@Test
	public void toStringTest() {
		assertEquals(DNA1.toString(), "[0.34,0.33,0.42,0.83,0.50,0.50]");
		assertEquals(OFB_DNA.toString(), "[1.00,0.33,0.42,0.00,0.00,0.00]");
		assertTrue(OFB_DNA.getGenes().stream().allMatch(gene -> gene >= 0 && gene <= 1));
	}

	@Test
	public void randomDNATest() {
		DNA dna = new DNA();
		assertEquals(dna.length(), 4);
		assertTrue(dna.getGenes().stream().allMatch(gene -> gene >= 0 && gene <= 1));
	}

	@Test
	public void crossOverTest() {
		DNA child1 = DNA1.crossover(DNA2);
		DNA child2 = DNA1.crossover(SHORT_DNA);
		DNA child3 = SHORT_DNA.crossover(DNA1);
		System.out.println("Normal x normal : " + child1);
		System.out.println("Normal x short : " + child2);
		System.out.println("Short x normal : " + child3);
		assertEquals(child1.length(), 6);
		assertTrue(child1.toString().equals("[0.13,0.65,0.45,0.83,0.50,0.50]") || child1.toString().equals("[0.34,0.33,0.42,0.84,0.52,0.95]"));
		assertEquals(child2.length(), 6);
		assertTrue(child2.toString().equals("[0.34,0.33,0.42,0.83,0.50,0.50]") || child2.toString().equals("[0.23,0.74,0.42,0.83,0.50,0.50]"));
		assertEquals(child3.length(), 2);
		assertTrue(child3.toString().equals("[0.23,0.33]") || child3.toString().equals("[0.34,0.74]"));
		
	}
	
	@Test
	public void mutateTest() {
		DNA.mutation_rate = 1;
		MUT_DNA1.mutate();
		assertNotEquals(MUT_DNA1.toString(), DNA1.toString());
		DNA.mutation_rate = 0;
		MUT_DNA2.mutate();
		assertEquals(MUT_DNA2.toString(), DNA1.toString());
	}
	
}
