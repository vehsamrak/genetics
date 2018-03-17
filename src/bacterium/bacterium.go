package bacterium

import "github.com/vehsamrak/genetics/src/applicationError"

type bacterium struct {
	lifePoints int
}

func New() bacterium {
	return bacterium{}
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

func (bacterium *bacterium) Move() (error error) {
	if !bacterium.IsAlive() {
		error = new(applicationError.CanNotMove)
	}

	return error
}
