package gene

type GenePool struct {
	cursor int
	genes  []gene
}

func (genePool *GenePool) Next() {
	genePoolLength := len(genePool.genes)

	if genePoolLength == 0 {
		return
	}

	if genePool.cursor < genePoolLength-1 {
		genePool.cursor++
	} else {
		genePool.cursor = 0
	}
}

func (genePool *GenePool) Add(gene gene) {
	genePool.genes = append(genePool.genes, gene)
}

func (genePool *GenePool) ExecuteCurrentGene() (ok bool, err error) {
	gene := genePool.genes[genePool.cursor]
	ok, err = gene.Act()

	return
}

func (genePool *GenePool) CountGenes() int {
	return len(genePool.genes)
}
