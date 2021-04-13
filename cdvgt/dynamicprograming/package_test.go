package dynamicprograming_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DynamicProgrammingTestSuit struct {
	suite.Suite
}

func TestDynamicProgramming(t *testing.T) {
	suite.Run(t, &DynamicProgrammingTestSuit{})
}
