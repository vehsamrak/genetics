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

func (suite *CellTest) TestGetLifePoints_givenNewCell_defaultAmountOfLifePointsReturned() {
	cell := New()

	lifePoints := cell.GetLifePoints()

	assert.Equal(suite.T(), 0, lifePoints)
}
