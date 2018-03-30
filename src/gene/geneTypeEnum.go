package gene

// GeneType represents type of gene
type GeneType int

const (
	GeneTypePhotosynthesize GeneType = iota
	GeneTypeEat
	GeneTypeMoveNorth
	GeneTypeMoveEast
	GeneTypeMoveSouth
	GeneTypeMoveWest
	GeneTypeWait
)
