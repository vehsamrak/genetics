package gene

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestPhotosynthesizeGene(test *testing.T) {
	suite.Run(test, &PhotosynthesizeGeneTestSuite{})
}

type PhotosynthesizeGeneTestSuite struct {
	suite.Suite
}

func (suite *PhotosynthesizeGeneTestSuite) Test_Code_noParameters_codeIsCorrect() {
	gene := &photosynthesizeGene{}

	assert.Equal(suite.T(), "P", gene.Code())
}
