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
	cell.lifePoints += points
}
