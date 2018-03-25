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

func (suite *gameFieldTestSuite) Test_New_noGameField_gameFieldCreatedWithDefaultValues() {
	gameField := New()

	assert.Len(suite.T(), gameField.microorganisms, bacteriumInitialCount)
}

func (suite *gameFieldTestSuite) Test_populate_emptyGameField_microorganismsAddedToGameField() {
	gameField := gameField{}

	gameField.populate()

	assert.Len(suite.T(), gameField.microorganisms, bacteriumInitialCount)
}
