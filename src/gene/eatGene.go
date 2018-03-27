package gene

type eatGene struct{}

func (eatGene *eatGene) Act() (ok bool, err error) {
	return true, nil
}

func (eatGene *eatGene) Code() string {
	return "F"
}
