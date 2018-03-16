package cell

type cell struct {
	lifePoints int
}

func New() cell {
	return cell{}
}

func (cell *cell) GetLifePoints() int {
	return cell.lifePoints
}

func (cell *cell) AddLifePoints(points int) {
	if points <= 0 {
		return
	}

	cell.lifePoints += points
}

func (cell *cell) DeductLifePoints(points int) {
	if points <= 0 {
		return
	}

	if cell.lifePoints - points < 0 {
		cell.lifePoints = 0
	} else {
		cell.lifePoints -= points
	}
}

func (cell *cell) IsAlive() bool {
	return cell.lifePoints > 0
}
