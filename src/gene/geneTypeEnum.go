package gene

// GeneType represents type of gene
type GeneType int

const (
	// GeneTypePhotosynthesize is gene for photosynthesis
	GeneTypePhotosynthesize GeneType = iota
	// GeneTypeEat is gene for eating
	GeneTypeEat
	// GeneTypeMoveNorth is gene for moving north
	GeneTypeMoveNorth
	// GeneTypeMoveEast is gene for moving east
	GeneTypeMoveEast
	// GeneTypeMoveSouth is gene for moving south
	GeneTypeMoveSouth
	// GeneTypeMoveWest is gene for moving west
	GeneTypeMoveWest
	// GeneTypeWait is gene for standing still
	GeneTypeWait
	genesCount
)

func GetAllGeneTypes() []GeneType {
	ts := make([]GeneType, genesCount)
	for i := 0; i < int(genesCount); i++ {
		ts[i] = GeneType(i)
	}

	return ts
}
