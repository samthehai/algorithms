package backtracking_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BackTrackingTestSuit struct {
	suite.Suite
}

func TestBackTracking(t *testing.T) {
	suite.Run(t, &BackTrackingTestSuit{})
}
