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
