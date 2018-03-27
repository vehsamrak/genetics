package gene

type moveSouthGene struct{}

func (moveSouthGene *moveSouthGene) Act() (ok bool, err error) {
	return true, nil
}

func (moveSouthGene *moveSouthGene) Code() string {
	return "v"
}
