package gene

type photosynthesizeGene struct{}

func (photosynthesizeGene *photosynthesizeGene) Act() (ok bool, err error) {
	return true, nil
}

func (photosynthesizeGene *photosynthesizeGene) Code() string {
	return "P"
}
