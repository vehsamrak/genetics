package gene

type moveEastGene struct{}

func (moveEastGene *moveEastGene) Act() (ok bool, err error) {
	return true, nil
}

func (moveEastGene *moveEastGene) Code() string {
	return ">"
}
