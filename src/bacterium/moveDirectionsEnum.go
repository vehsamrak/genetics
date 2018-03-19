package bacterium

// Direction of bacterium movement based on quadrille (2d map)
type Direction int

const (
	directionNorth Direction = iota
	directionEast
	directionSouth
	directionWest
)
