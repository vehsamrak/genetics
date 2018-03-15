package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CellTest struct {
	suite.Suite
}

func TestCell(test *testing.T) {
	suite.Run(test, new(CellTest))
}

func (suite *CellTest) TestGetLifePoints_newCell_defaultAmountOfLifePointsReturned() {
	cell := New()

	lifePoints := cell.GetLifePoints()

	assert.Equal(suite.T(), 0, lifePoints)
}

func (suite *CellTest) TestAddLifePoints_newCell_cellContainsAmountOfLifePoints() {
	cell := New()
	additionalLifePoints := 1

	cell.AddLifePoints(additionalLifePoints)
	lifePoints := cell.GetLifePoints()

	assert.Equal(suite.T(), additionalLifePoints, lifePoints)
}
