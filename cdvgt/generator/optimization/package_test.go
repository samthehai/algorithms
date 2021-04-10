package optimization_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OptimizationTestSuit struct {
	suite.Suite
}

func TestOptimization(t *testing.T) {
	suite.Run(t, &OptimizationTestSuit{})
}
