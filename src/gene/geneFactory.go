package gene

// GeneFactory is a point to creates genes
type GeneFactory struct {
}

// Create new gene by its type
func (factory *GeneFactory) Create(geneType GeneType) (gene gene) {
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
