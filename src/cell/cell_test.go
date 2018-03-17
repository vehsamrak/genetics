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
	deduction                    int
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

var additionTestTable = []struct {
	cellLifePoints              int
	addition                    int
	cellLifePointsAfterAddition int
}{
	{0, 0, 0},
	{0, 1, 1},
	{1, 1, 2},
	{1, 2, 3},
	{2, 3, 5},
	{1, -1, 1},
	{2, -1, 2},
}

var isAliveTestTable = []struct {
	cellLifePoints int
	isAlive        bool
}{
	{0, false},
	{1, true},
	{-1, false},
}

func (suite *CellTest) Test_GetLifePoints_newCell_defaultAmountOfLifePointsReturned() {
	cell := New()
	defaultLifePoints := 0

	lifePoints := cell.GetLifePoints()

	assert.Equal(suite.T(), defaultLifePoints, lifePoints)
}

func (suite *CellTest) Test_AddLifePoints_cellWithPoints_cellContainsPoints() {
	for id, dataset := range additionTestTable {
		cell := cell{lifePoints: dataset.cellLifePoints}

		cell.AddLifePoints(dataset.addition)

		assert.Equal(suite.T(), dataset.cellLifePointsAfterAddition, cell.lifePoints, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *CellTest) Test_DeductLifePoints_cellWithPoints_cellContainsPoints() {
	for id, dataset := range deductionTestTable {
		cell := cell{lifePoints: dataset.cellLifePoints}

		cell.DeductLifePoints(dataset.deduction)

		assert.Equal(suite.T(), dataset.cellLifePointsAfterDeduction, cell.lifePoints, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *CellTest) Test_IsAlive_cellWithPoints_cellAliveStatusReturned() {
	for id, dataset := range isAliveTestTable {
		cell := cell{lifePoints: dataset.cellLifePoints}

		isCellAlive := cell.IsAlive()

		assert.Equal(suite.T(), dataset.isAlive, isCellAlive, fmt.Sprintf("Dataset #%v", id))
	}
}
