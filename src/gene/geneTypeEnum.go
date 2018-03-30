package gene

// GeneType represents type of gene
type GeneType int

const (
	// GeneTypePhotosynthesize: gene for photosynthesis
	GeneTypePhotosynthesize GeneType = iota
	// GeneTypeEat: gene for eating
	GeneTypeEat
	// GeneTypeMoveNorth: gene for moving north
	GeneTypeMoveNorth
	// GeneTypeMoveEast: gene for moving east
	GeneTypeMoveEast
	// GeneTypeMoveSouth: gene for moving south
	GeneTypeMoveSouth
	// GeneTypeMoveWest: gene for moving west
	GeneTypeMoveWest
	// GeneTypeWait: gene for standing still
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
