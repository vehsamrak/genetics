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
	genesCount
)

func getAllGeneTypes() []GeneType {
	ts := make([]GeneType, genesCount)
	for i := 0; i < int(genesCount); i++ {
		ts[i] = GeneType(i)
	}

	return ts
}
