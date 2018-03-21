package bacterium

type microorganism interface {
	IsAlive() bool
	IsDead() bool
	X() int
	Y() int
}
