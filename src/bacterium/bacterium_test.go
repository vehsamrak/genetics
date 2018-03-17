package bacterium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

var isAliveTestTable = []struct {
	lifePoints int
	isAlive    bool
}{
	{0, false},
	{1, true},
	{-1, false},
}

func (suite *BacteriumTest) Test_GetLifePoints_newBacterium_defaultAmountOfLifePointsReturned() {
	bacterium := New()
	defaultLifePoints := 0

	lifePoints := bacterium.GetLifePoints()

	assert.Equal(suite.T(), defaultLifePoints, lifePoints)
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
