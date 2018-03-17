package bacterium

import "github.com/vehsamrak/genetics/src/applicationError"

const DEFAULT_LIFE_POINTS = 10

type bacterium struct {
	lifePoints int
	X          int
	Y          int
}

func New() bacterium {
	return bacterium{lifePoints: DEFAULT_LIFE_POINTS}
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

	if bacterium.lifePoints - points < 0 {
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
	case DIRECTION_NORTH:
		bacterium.Y++
	case DIRECTION_EAST:
		bacterium.X++
	case DIRECTION_SOUTH:
		bacterium.Y--
	case DIRECTION_WEST:
		bacterium.X--
	}

	bacterium.lifePoints--

	return error
}
