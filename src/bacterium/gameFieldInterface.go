package bacterium

type gameField interface {
	allBacterias() []microorganism
	addBacterium(microorganism microorganism)
}
