package gene

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestGenePool(test *testing.T) {
	suite.Run(test, &GenePoolTestSuite{})
}

type GenePoolTestSuite struct {
	suite.Suite
}

var nextTestTable = []struct {
	genesCount     int
	initialCursor  int
	expectedCursor int
}{
	{0, 0, 0},
	{2, 1, 0},
	{2, 0, 1},
	{3, 1, 2},
}

func (suite *GenePoolTestSuite) Test_Next_genePoolWithGenesCollectionAndCursor_cursorMovedAsExpected() {
	for id, dataset := range nextTestTable {
		genePool := genePool{cursor: dataset.initialCursor, genes: make([]gene, dataset.genesCount)}

		genePool.Next()

		assert.Equal(suite.T(), dataset.expectedCursor, genePool.cursor, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *GenePoolTestSuite) Test_Add_genePoolWithoutGenes_geneAddedToPool() {
	genePool := genePool{}
	gene := &photosynthesizeGene{}

	genePool.Add(gene)

	assert.Len(suite.T(), genePool.genes, 1)
}

func (suite *GenePoolTestSuite) Test_ExecuteCurrentGene_genePoolWithGeneAndCursorOnIt_geneUnderCursorActs() {
	genePool := genePool{}
	genePool.Add(&photosynthesizeGene{})
	genePool.Add(&moveNorthGene{})

	for range genePool.genes {
		genePool.Next()
		ok, err := genePool.ExecuteCurrentGene()

		assert.True(suite.T(), ok)
		assert.Nil(suite.T(), err)
	}

}
