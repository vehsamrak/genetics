package gene

import (
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

func (suite *GenePoolTestSuite) Test_Next_genePoolWithTwoGenesAndCursorOnZero_cursorMovedToNextGene() {
	genes := make([]*gene, 2)
	genePool := genePool{cursor: 0, genes: genes}

	genePool.Next()

	assert.Equal(suite.T(), 1, genePool.cursor)
}

func (suite *GenePoolTestSuite) Test_Next_genePoolWithoutGenesAndCursorOnZero_cursorStaysOnZero() {
	genePool := genePool{cursor: 0}

	genePool.Next()

	assert.Equal(suite.T(), 0, genePool.cursor)
}
