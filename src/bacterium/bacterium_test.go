package bacterium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/genetics/src/applicationError"
	Gene "github.com/vehsamrak/genetics/src/gene"
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
	{directionNorth, 0, -1},
	{directionEast, 1, 0},
	{directionSouth, 0, 1},
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

	assert.Equal(suite.T(), bacterium.LifePoints(), defaultLifePoints)
}

func (suite *BacteriumTestSuite) Test_New_newBacterium_bacteriumContainsNonEmptyGenePool() {
	geneTypesCount := len(Gene.GetAllGeneTypes())

	bacterium := New()
	genePool := bacterium.GenePool()
	bacteriumGenesCount := genePool.CountGenes()

	assert.Equal(suite.T(), geneTypesCount, bacteriumGenesCount)
}

func (suite *BacteriumTestSuite) Test_GetLifePoints_bacteriumWithOneLifePoint_correctAmountOfLifePointsReturned() {
	bacterium := bacterium{lifePoints: 1}

	lifePoints := bacterium.LifePoints()

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
	bacterium := bacterium{lifePoints: 0, gameField: suite.createGameField()}

	err := bacterium.Move(directionNorth)

	_, ok := err.(*applicationError.CanNotMove)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "Microorganism can't move", err.Error())
}

func (suite *BacteriumTestSuite) Test_Move_bacteriumWithTwoLifePoints_bacteriumMovedAndOneLifePointLeft() {
	for id, dataset := range moveTestTable {
		bacterium := bacterium{lifePoints: 2, gameField: suite.createGameField()}

		err := bacterium.Move(dataset.direction)

		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), dataset.X, bacterium.x, fmt.Sprintf("Dataset #%v", id))
		assert.Equal(suite.T(), dataset.Y, bacterium.y, fmt.Sprintf("Dataset #%v", id))
		assert.Equal(suite.T(), 1, bacterium.lifePoints, fmt.Sprintf("Dataset #%v", id))
	}
}

func (suite *BacteriumTestSuite) Test_Move_bacteriumWithOneLifePoint_bacteriumMovedAndBecameDead() {
	bacterium := bacterium{lifePoints: 1, gameField: suite.createGameField()}

	err := bacterium.Move(directionNorth)

	assert.Nil(suite.T(), err)
	assert.True(suite.T(), bacterium.IsDead())
}

func (suite *BacteriumTestSuite) Test_Move_bacteriumMovesToAnotherOne_stuckErrorAndFirstBacteriumLostLifePoint() {
	gameField := suite.createGameFieldWithBacterium(1, 0, 0)
	bacterium := &bacterium{lifePoints: 1, x: 0, y: 1, gameField: gameField}
	gameField.addBacterium(bacterium)

	err := bacterium.Move(directionNorth)

	assert.NotNil(suite.T(), err)
	_, ok := err.(*applicationError.Stuck)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "Microorganism can't move to the field where another one stands", err.Error())
	assert.Equal(suite.T(), 0, bacterium.lifePoints)
}

func (suite *BacteriumTestSuite) Test_Eat_bacteriumWithOneLifePointEatsDeadOne_bacteriumLifePointsIncreasedAndDeadBacteriumDisappears() {
	gameField := suite.createGameFieldWithDeadBacteriumInZeroCell()
	initialLifePoints := 1
	lifePointsAfterMeal := initialLifePoints - lifePointsCostEat + lifePointsGainEat
	bacterium := &bacterium{lifePoints: initialLifePoints, x: 0, y: 1, gameField: gameField}
	gameField.addBacterium(bacterium)

	bacterium.Eat(directionNorth)

	assert.Equal(suite.T(), lifePointsAfterMeal, bacterium.lifePoints)
	assert.Len(suite.T(), gameField.allBacterias(), 1)
}

func (suite *BacteriumTestSuite) Test_Eat_bacteriumWithOneLifePointAndNoMoreBacterias_bacteriumLifePointsDecreasedToZero() {
	bacterium := &bacterium{lifePoints: 1, gameField: suite.createGameField()}

	bacterium.Eat(directionNorth)

	assert.Equal(suite.T(), 0, bacterium.lifePoints)
}

func (suite *BacteriumTestSuite) Test_Eat_deadBacterium_isDeadError() {
	bacterium := &bacterium{lifePoints: 0, gameField: suite.createGameField()}

	err := bacterium.Eat(directionNorth)

	_, ok := err.(*applicationError.IsDead)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "Microorganism can't process actions after death", err.Error())
}

func (suite *BacteriumTestSuite) Test_Photosynthesize_aliveBacterium_lifePointsIncreased() {
	lifePointsBeforePhotosynthesize := 1
	bacterium := &bacterium{lifePoints: lifePointsBeforePhotosynthesize}
	lifePointsAfterPhotosynthesize := lifePointsBeforePhotosynthesize + lifePointsGainPhotosynthesize

	bacterium.Photosynthesize()

	assert.Equal(suite.T(), lifePointsAfterPhotosynthesize, bacterium.lifePoints)
}

func (suite *BacteriumTestSuite) Test_GenePool_bacteriumWithGenes_genePoolReturned() {
	bacterium := suite.createBacteriumWithGenes()

	genePool := bacterium.GenePool()

	assert.Equal(suite.T(), 1, genePool.CountGenes())
}

func (suite *BacteriumTestSuite) createBacteriumWithGenes() *bacterium {
	geneFactory := Gene.GeneFactory{}
	gene := geneFactory.CreateGene(Gene.GeneTypeWait)

	genePool := &Gene.GenePool{}
	genePool.Add(gene)

	bacterium := &bacterium{genePool: genePool}

	return bacterium
}

func (suite *BacteriumTestSuite) Test_Photosynthesize_deadBacterium_isDeadError() {
	bacterium := &bacterium{lifePoints: 0}

	err := bacterium.Photosynthesize()

	_, ok := err.(*applicationError.IsDead)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "Microorganism can't process actions after death", err.Error())
}

func (suite *BacteriumTestSuite) createGameFieldWithDeadBacteriumInZeroCell() gameField {
	return suite.createGameFieldWithBacterium(0, 0, 0)
}

func (suite *BacteriumTestSuite) createGameFieldWithBacterium(lifePoints int, x int, y int) gameField {
	gameField := suite.createGameField()
	bacterium := &bacterium{lifePoints: lifePoints, x: x, y: y, gameField: gameField}
	gameField.addBacterium(bacterium)

	return gameField
}

func (suite *BacteriumTestSuite) createGameField() gameField {
	return &gameFieldMock{}
}

type gameFieldMock struct {
	bacterias []microorganism
}

func (field *gameFieldMock) allBacterias() []microorganism {
	return field.bacterias
}

func (field *gameFieldMock) addBacterium(microorganism microorganism) {
	field.bacterias = append(field.bacterias, microorganism)
}

func (field *gameFieldMock) removeBacterium(microorganism microorganism) {
	for bacteriumKey, bacterium := range field.bacterias {
		if bacterium == microorganism {
			field.bacterias = append(field.bacterias[:bacteriumKey], field.bacterias[bacteriumKey+1:]...)
			break
		}
	}
}
