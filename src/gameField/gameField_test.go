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

func (suite *gameFieldTestSuite) Test_Populate_emptyGameField_microorganismsAddedToGameField() {
	gameField := gameField{}

	gameField.Populate()

	assert.Len(suite.T(), gameField.microorganisms, microorganismsInitialCount)
}
