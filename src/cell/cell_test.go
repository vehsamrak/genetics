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
	defaultLifePoints := 0

	lifePoints := cell.GetLifePoints()

	assert.Equal(suite.T(), defaultLifePoints, lifePoints)
}

func (suite *CellTest) TestAddLifePoints_addOnePointToCellWithZeroPoints_pointsAdded() {
	cell := cell{lifePoints: 0}
	additionalLifePoints := 1

	cell.AddLifePoints(additionalLifePoints)

	assert.Equal(suite.T(), additionalLifePoints, cell.GetLifePoints())
}

func (suite *CellTest) TestDeductLifePoints_cellWithZeroPoints_cellContainsZeroPoints() {
	cell := cell{lifePoints: 0}
	lifePointsToDeduct := 1

	cell.DeductLifePoints(lifePointsToDeduct)

	assert.Equal(suite.T(), 0, cell.GetLifePoints())
}

func (suite *CellTest) TestDeductLifePoints_cellWithOnePoint_cellContainsZeroPoints() {
	cell := cell{lifePoints: 1}
	lifePointsToDeduct := 1

	cell.DeductLifePoints(lifePointsToDeduct)

	assert.Equal(suite.T(), 0, cell.GetLifePoints())
}

func (suite *CellTest) TestDeductLifePoints_deductOnePointFromCellWithTwoPoints_cellContainsOnePoint() {
	cell := cell{lifePoints: 2}
	lifePointsToDeduct := 1

	cell.DeductLifePoints(lifePointsToDeduct)

	assert.Equal(suite.T(), 1, cell.GetLifePoints())
}
