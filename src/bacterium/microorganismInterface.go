package bacterium

type microorganism interface {
	GameField() gameField
	IsAlive() bool
	IsDead() bool
	X() int
	Y() int
	LifePoints() int
}
