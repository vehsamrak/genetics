package bacterium

// Bacterium move direction based on quadrille (2d map)
type MoveDirection int

const (
	directionNorth MoveDirection = iota
	directionEast
	directionSouth
	directionWest
)
