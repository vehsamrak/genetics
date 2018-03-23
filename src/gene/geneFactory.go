package gene

type GeneFactory struct {
}

func (factory *GeneFactory) Create(geneType GeneType) (gene gene) {
	switch geneType {
	case geneTypePhotosynthesize:
		gene = &photosynthesizeGene{}
	}

	return
}
