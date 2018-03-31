package bacterium

import "github.com/vehsamrak/genetics/src/gene"

type microorganism interface {
	GameField() gameField
	IsAlive() bool
	GenePool() *gene.GenePool
	IsDead() bool
	X() int
	Y() int
	LifePoints() int
}
