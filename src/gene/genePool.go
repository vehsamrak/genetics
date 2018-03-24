package gene

type genePool struct {
	cursor int
	genes  []*gene
}

func (genePool *genePool) Next() {
	if len(genePool.genes) > 0 {
		genePool.cursor++
	}
}
