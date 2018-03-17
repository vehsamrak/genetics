package bacterium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/genetics/src/applicationError"
)

type BacteriumTest struct {
	suite.Suite
}

func TestBacterium(test *testing.T) {
	suite.Run(test, new(BacteriumTest))
}

var deductionTestTable = []struct {
	lifePoints               int
	deduction                int
	lifePointsAfterDeduction int
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
	lifePoints              int
	addition                int
	lifePointsAfterAddition int
}{
	{0, 0, 0},
	{0, 1, 1},
	{1, 1, 2},
	{1, 2, 3},
	{2, 3, 5},
	{1, -1, 1},
	{2, -1, 2},
}

var moveTestTable = []struct {
	direction MoveDirection
	X         int
	Y         int
}{
	{DIRECTION_NORTH, 0, 1},
	{DIRECTION_EAST, 1, 0},
	{DIRECTION_SOUTH, 0, -1},
	{DIRECTION_WEST, -1, 0},
}

var isAliveTestTable = []struct {
	lifePoints int
	isAlive    bool
}{
	{0, false},
	{1, true},
	{-1, false},
}

func (suite *BacteriumTest) Test_New_newBacterium_defaultAmountOfLifePoints() {
	bacterium := New()
	defaultLifePoints := 10

	assert.Equal(suite.T(), bacterium.lifePoints, defaultLifePoints)
}

func (suite *BacteriumTest) Test_GetLifePoints_bacteriumWithOneLifePoint_correctAmountOfLifePointsReturned() {
	bacterium := bacterium{lifePoints: 1}

	lifePoints := bacterium.GetLifePoints()

	assert.Equal(suite.T(), 1, lifePoints)
}

func (suite *BacteriumTest) Test_AddLifePoints_bacteriumWithPoints_bacteriumContainsPoints() {
	for id, dataset := range additionTestTable {
		bacterium := bacterium{lifePoints: dataset.lifePoints}

		bacterium.AddLifePoints(dataset.addition)

		assert.Equal(suite.T(), dataset.lifePointsAfterAddition, bacterium.lifePoints, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *BacteriumTest) Test_DeductLifePoints_bacteriumWithPoints_bacteriumContainsPoints() {
	for id, dataset := range deductionTestTable {
		bacterium := bacterium{lifePoints: dataset.lifePoints}

		bacterium.DeductLifePoints(dataset.deduction)

		assert.Equal(suite.T(), dataset.lifePointsAfterDeduction, bacterium.lifePoints, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *BacteriumTest) Test_IsAlive_bacteriumWithPoints_bacteriumAliveStatusReturned() {
	for id, dataset := range isAliveTestTable {
		bacterium := bacterium{lifePoints: dataset.lifePoints}

		isBacteriumAlive := bacterium.IsAlive()

		assert.Equal(suite.T(), dataset.isAlive, isBacteriumAlive, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *BacteriumTest) Test_Move_bacteriumWithoutLifePoints_canNotMoveError() {
	bacterium := bacterium{lifePoints: 0}

	err := bacterium.Move(DIRECTION_NORTH)

	_, ok := err.(*applicationError.CanNotMove)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "Bacterium can't move", err.Error())
}

func (suite *BacteriumTest) Test_Move_bacteriumWithTwoLifePoints_bacteriumMovedAndOneLifePointLeft() {
	for id, dataset := range moveTestTable {
		bacterium := bacterium{lifePoints: 2}

		err := bacterium.Move(dataset.direction)

		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), dataset.X, bacterium.X, fmt.Sprintf("Dataset #%v", id))
		assert.Equal(suite.T(), 1, bacterium.lifePoints, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *BacteriumTest) Test_Move_bacteriumWithOneLifePoint_bacteriumMovedAndBecameDead() {
	bacterium := bacterium{lifePoints: 1}

	err := bacterium.Move(DIRECTION_NORTH)

	assert.Nil(suite.T(), err)
	assert.True(suite.T(), bacterium.IsDead())
}
