package gene

type genePool struct {
	cursor int
	genes  []*gene
}

func (genePool *genePool) Next() {
	genePool.cursor++
}
