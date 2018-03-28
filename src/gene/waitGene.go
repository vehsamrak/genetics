package gene

type waitGene struct{}

func (waitGene *waitGene) Act() (ok bool, err error) {
	return true, nil
}

func (waitGene *waitGene) Code() string {
	return "."
}
