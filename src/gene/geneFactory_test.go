package gene

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestGeneFactory(test *testing.T) {
	suite.Run(test, &GeneFactoryTestSuite{})
}

type GeneFactoryTestSuite struct {
	suite.Suite
}

func (suite *GeneFactoryTestSuite) Test_Code_photosynthesizeGeneTypeCreation_newGeneFactoryCreatedAndCodeIsCorrect() {
	factory := &GeneFactory{}

	gene := factory.Create(geneTypePhotosynthesize)

	_, ok := gene.(*photosynthesizeGene)
	assert.True(suite.T(), ok)
}
