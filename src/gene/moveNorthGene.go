package gene

type moveNorthGene struct{}

func (moveNorthGene *moveNorthGene) Act() (ok bool, err error) {
	return true, nil
}

func (moveNorthGene *moveNorthGene) Code() string {
	return "^"
}
