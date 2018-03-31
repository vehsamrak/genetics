package gene

// GeneFactory is a point to creates genes
type GeneFactory struct {
}

// CreateGene creates new gene by its type
func (factory *GeneFactory) CreateGene(geneType GeneType) (gene gene) {
	switch geneType {
	case GeneTypeEat:
		gene = &eatGene{}
	case GeneTypeMoveEast:
		gene = &moveEastGene{}
	case GeneTypeMoveNorth:
		gene = &moveNorthGene{}
	case GeneTypeMoveSouth:
		gene = &moveSouthGene{}
	case GeneTypeMoveWest:
		gene = &moveWestGene{}
	case GeneTypePhotosynthesize:
		gene = &photosynthesizeGene{}
	case GeneTypeWait:
		gene = &waitGene{}
	}

	return
}

// CreateGenePool creates new gene pool with genes in it
func (factory *GeneFactory) CreateGenePool() (genePool *GenePool) {
	genePool = &GenePool{}
	var geneType GeneType

	for _, geneType = range GetAllGeneTypes() {
		gene := factory.CreateGene(geneType)
		genePool.Add(gene)
	}

	return
}
