package bacterium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/genetics/src/applicationError"
)

func TestBacterium(test *testing.T) {
	suite.Run(test, new(BacteriumTestSuite))
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
	direction Direction
	X         int
	Y         int
}{
	{directionNorth, 0, 1},
	{directionEast, 1, 0},
	{directionSouth, 0, -1},
	{directionWest, -1, 0},
}

var isAliveTestTable = []struct {
	lifePoints int
	isAlive    bool
}{
	{0, false},
	{1, true},
	{-1, false},
}

type BacteriumTestSuite struct {
	suite.Suite
}

func (suite *BacteriumTestSuite) Test_New_newBacterium_defaultAmountOfLifePoints() {
	bacterium := New()
	defaultLifePoints := 10

	assert.Equal(suite.T(), bacterium.lifePoints, defaultLifePoints)
}

func (suite *BacteriumTestSuite) Test_GetLifePoints_bacteriumWithOneLifePoint_correctAmountOfLifePointsReturned() {
	bacterium := bacterium{lifePoints: 1}

	lifePoints := bacterium.GetLifePoints()

	assert.Equal(suite.T(), 1, lifePoints)
}

func (suite *BacteriumTestSuite) Test_AddLifePoints_bacteriumWithPoints_bacteriumContainsPoints() {
	for id, dataset := range additionTestTable {
		bacterium := bacterium{lifePoints: dataset.lifePoints}

		bacterium.AddLifePoints(dataset.addition)

		assert.Equal(suite.T(), dataset.lifePointsAfterAddition, bacterium.lifePoints, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *BacteriumTestSuite) Test_DeductLifePoints_bacteriumWithPoints_bacteriumContainsPoints() {
	for id, dataset := range deductionTestTable {
		bacterium := bacterium{lifePoints: dataset.lifePoints}

		bacterium.DeductLifePoints(dataset.deduction)

		assert.Equal(suite.T(), dataset.lifePointsAfterDeduction, bacterium.lifePoints, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *BacteriumTestSuite) Test_IsAlive_bacteriumWithPoints_bacteriumAliveStatusReturned() {
	for id, dataset := range isAliveTestTable {
		bacterium := bacterium{lifePoints: dataset.lifePoints}

		isBacteriumAlive := bacterium.IsAlive()

		assert.Equal(suite.T(), dataset.isAlive, isBacteriumAlive, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *BacteriumTestSuite) Test_Move_bacteriumWithoutLifePoints_canNotMoveError() {
	bacterium := bacterium{lifePoints: 0}

	err := bacterium.Move(directionNorth)

	_, ok := err.(*applicationError.CanNotMove)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "Bacterium can't move", err.Error())
}

func (suite *BacteriumTestSuite) Test_Move_bacteriumWithTwoLifePoints_bacteriumMovedAndOneLifePointLeft() {
	for id, dataset := range moveTestTable {
		bacterium := bacterium{lifePoints: 2}

		err := bacterium.Move(dataset.direction)

		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), dataset.X, bacterium.X, fmt.Sprintf("Dataset #%v", id))
		assert.Equal(suite.T(), 1, bacterium.lifePoints, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *BacteriumTestSuite) Test_Move_bacteriumWithOneLifePoint_bacteriumMovedAndBecameDead() {
	bacterium := bacterium{lifePoints: 1}

	err := bacterium.Move(directionNorth)

	assert.Nil(suite.T(), err)
	assert.True(suite.T(), bacterium.IsDead())
}

func (suite *BacteriumTestSuite) Test_Move_bacteriumNearAnotherOne_cantMoveToCoordinatesWhereAnotherStandsError() {
	gameField := createGameField()
	_ = bacterium{lifePoints: defaultLifePoints, X: 0, Y: 1, gameField: &gameField}
	southBacterium := bacterium{lifePoints: defaultLifePoints, X: 0, Y: 0, gameField: &gameField}

	err := southBacterium.Move(directionNorth)

	assert.Nil(suite.T(), err)
}

func createGameField() gameField {
	mock := &gameFieldMock{}

	return mock
}

type gameFieldMock struct {
}

func (field *gameFieldMock) getAllBacterias() {
}
