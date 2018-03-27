package gene

type moveWestGene struct{}

func (moveWestGene *moveWestGene) Act() (ok bool, err error) {
	return true, nil
}

func (moveWestGene *moveWestGene) Code() string {
	return "<"
}
