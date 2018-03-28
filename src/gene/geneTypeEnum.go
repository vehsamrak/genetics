package gene

// GeneType represents type of gene
type GeneType int

const (
	geneTypePhotosynthesize GeneType = iota
	geneTypeEat
	geneTypeMoveNorth
	geneTypeMoveEast
	geneTypeMoveSouth
	geneTypeMoveWest
	geneTypeWait
)
