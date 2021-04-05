package generation_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GenerationTestSuit struct {
	suite.Suite
}

func TestGeneration(t *testing.T) {
	suite.Run(t, &GenerationTestSuit{})
}
