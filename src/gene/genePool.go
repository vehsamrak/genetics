package gene

type genePool struct {
	cursor int
	genes  []gene
}

func (genePool *genePool) Next() {
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

func (genePool *genePool) Add(gene gene) {
	genePool.genes = append(genePool.genes, gene)
}

func (genePool *genePool) ExecuteCurrentGene() (ok bool, err error) {
	gene := genePool.genes[genePool.cursor]
	ok, err = gene.Act()

	return
}
