package gameField

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestGameField(test *testing.T) {
	suite.Run(test, &gameFieldTestSuite{})
}

type gameFieldTestSuite struct {
	suite.Suite
}

func (suite *gameFieldTestSuite) Test_New_defaultBacteriumInitialCount_gameFieldCreatedWithExpectedAmountOfBacterias() {
	gameField := New()

	assert.Len(suite.T(), gameField.allBacterias(), bacteriumInitialCount)
}

func (suite *gameFieldTestSuite) Test_addBacterium_fieldWithoutBacterias_bacteriumAddedToField() {
	gameField := gameField{}
	microorganism := createMicroorganism()

	gameField.addBacterium(microorganism)

	assert.Len(suite.T(), gameField.microorganisms, 1)
}

func (suite *gameFieldTestSuite) Test_removeBacterium_fieldWithOneBacterium_bacteriumRemovedFromField() {
	gameField := gameField{}
	microorganism := createMicroorganism()
	gameField.addBacterium(microorganism)

	gameField.removeBacterium(microorganism)

	assert.Len(suite.T(), gameField.microorganisms, 0)
}

func (suite *gameFieldTestSuite) Test_populate_emptyGameField_microorganismsAddedToGameField() {
	gameField := gameField{}

	gameField.populate()

	assert.Len(suite.T(), gameField.microorganisms, bacteriumInitialCount)
}

func createMicroorganism() microorganism {
	return &bacteriumMock{}
}

type bacteriumMock struct {
}

func (bacterium *bacteriumMock) IsAlive() bool {
	return true
}
