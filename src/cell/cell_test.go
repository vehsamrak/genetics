package cell

import (
	"fmt"
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

var deductionTestTable = []struct {
	cellLifePoints               int
	deductionCount               int
	cellLifePointsAfterDeduction int
}{
	{0, 0, 0},
	{0, 1, 0},
	{1, 1, 0},
	{2, 1, 1},
	{3, 2, 1},
	{3, 3, 0},
	{3, 4, 0},
	{1, -1, 1},
	{2, -1, 2},
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

func (suite *CellTest) TestDeductLifePoints_cellWithPoints_cellContainsPoints() {
	for key, tt := range deductionTestTable {

		cell := cell{lifePoints: tt.cellLifePoints}
		cell.DeductLifePoints(tt.deductionCount)

		assert.Equal(suite.T(), tt.cellLifePointsAfterDeduction, cell.GetLifePoints(), fmt.Sprintf("Dataset #%v", key))
	}
}
