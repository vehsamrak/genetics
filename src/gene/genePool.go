package gene

// GenePool is gene collection
type GenePool struct {
	cursor int
	genes  []gene
}

// Next moves cursor to next gene in gene pool
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

// Add gene to gene pool
func (genePool *GenePool) Add(gene gene) {
	genePool.genes = append(genePool.genes, gene)
}

// ExecuteCurrentGene runs action of the gene that is currently pointed by cursor
func (genePool *GenePool) ExecuteCurrentGene() (ok bool, err error) {
	gene := genePool.genes[genePool.cursor]
	ok, err = gene.Act()

	return
}

// CountGenes for counting all genes in pool
func (genePool *GenePool) CountGenes() int {
	return len(genePool.genes)
}
