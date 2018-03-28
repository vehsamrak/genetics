package gene

// GeneFactory is a point to creates genes
type GeneFactory struct {
}

// Create new gene by its type
func (factory *GeneFactory) Create(geneType GeneType) (gene gene) {
	switch geneType {
	case geneTypeEat:
		gene = &eatGene{}
	case geneTypeMoveEast:
		gene = &moveEastGene{}
	case geneTypeMoveNorth:
		gene = &moveNorthGene{}
	case geneTypeMoveSouth:
		gene = &moveSouthGene{}
	case geneTypeMoveWest:
		gene = &moveWestGene{}
	case geneTypePhotosynthesize:
		gene = &photosynthesizeGene{}
	case geneTypeWait:
		gene = &waitGene{}
	}

	return
}
