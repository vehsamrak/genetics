package bacterium

import "github.com/vehsamrak/genetics/src/applicationError"

const defaultLifePoints = 10

type bacterium struct {
	gameField  *gameField
	lifePoints int
	X          int
	Y          int
}

// New bacterium constructor
func New() bacterium {
	return bacterium{lifePoints: defaultLifePoints}
}

func (bacterium *bacterium) GetLifePoints() int {
	return bacterium.lifePoints
}

func (bacterium *bacterium) AddLifePoints(points int) {
	if points <= 0 {
		return
	}

	bacterium.lifePoints += points
}

func (bacterium *bacterium) DeductLifePoints(points int) {
	if points <= 0 {
		return
	}

	if bacterium.lifePoints-points < 0 {
		bacterium.lifePoints = 0
	} else {
		bacterium.lifePoints -= points
	}
}

func (bacterium *bacterium) IsAlive() bool {
	return bacterium.lifePoints > 0
}

func (bacterium *bacterium) IsDead() bool {
	return bacterium.lifePoints <= 0
}

func (bacterium *bacterium) Move(direction MoveDirection) (error error) {
	if bacterium.IsDead() {
		error = new(applicationError.CanNotMove)
	}

	switch direction {
	case directionNorth:
		bacterium.Y++
	case directionEast:
		bacterium.X++
	case directionSouth:
		bacterium.Y--
	case directionWest:
		bacterium.X--
	}

	bacterium.lifePoints--

	return error
}
