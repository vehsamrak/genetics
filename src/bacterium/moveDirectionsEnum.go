package bacterium

// Move direction of bacterium based on quadrille (2d map)
type MoveDirection int

const (
	directionNorth MoveDirection = iota
	directionEast
	directionSouth
	directionWest
)
