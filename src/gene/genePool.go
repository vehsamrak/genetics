package gene

type genePool struct {
	cursor int
	genes  []*gene
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
