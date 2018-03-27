package gene

import (
	"fmt"
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

var geneTypeTestTable = []struct {
	geneType GeneType
	geneCode string
}{
	{geneTypePhotosynthesize, "P"},
	{geneTypeMoveNorth, "^"},
	{geneTypeMoveEast, ">"},
	{geneTypeMoveSouth, "v"},
	{geneTypeMoveWest, "<"},
	{geneTypeEat, "F"},
}

func (suite *GeneFactoryTestSuite) Test_Create_geneType_newGeneCreatedAndCodeIsCorrect() {
	factory := &GeneFactory{}

	for id, dataset := range geneTypeTestTable {
		gene := factory.Create(dataset.geneType)

		assert.Equal(suite.T(), dataset.geneCode, gene.Code(), fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *GeneFactoryTestSuite) Test_Code_moveNorthGeneTypeCreation_newGeneFactoryCreatedAndCodeIsCorrect() {
	factory := &GeneFactory{}

	gene := factory.Create(geneTypePhotosynthesize)

	_, ok := gene.(*photosynthesizeGene)
	assert.True(suite.T(), ok)
}
