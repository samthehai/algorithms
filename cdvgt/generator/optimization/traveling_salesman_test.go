package optimization_test

import (
	"testing"

	"github.com/samthehai/algorithms/cdvgt/generator/optimization"
)

func (s *OptimizationTestSuit) TestTravelingSalesman() {
	testcases := []struct {
		name      string
		n         uint
		startAt   uint
		distances []optimization.CityDistance
		expected  optimization.TravelRoute
	}{
		{
			name:      "n = 0",
			n:         0,
			distances: []optimization.CityDistance{},
			expected:  optimization.TravelRoute{},
		},
		{
			name:    "n = 6",
			n:       6,
			startAt: 1,
			distances: []optimization.CityDistance{
				{1, 2, 3},
				{1, 3, 2},
				{1, 4, 1},
				{2, 3, 1},
				{2, 4, 2},
				{3, 4, 4},
			},
			expected: optimization.TravelRoute{
				Route: []uint{1, 3, 2, 4, 1},
				Cost:  6,
			},
		},
		{
			name:    "n = 6",
			n:       6,
			startAt: 4,
			distances: []optimization.CityDistance{
				{1, 2, 3},
				{1, 3, 2},
				{1, 4, 1},
				{2, 3, 1},
				{2, 4, 2},
				{3, 4, 4},
			},
			expected: optimization.TravelRoute{
				Route: []uint{4, 1, 3, 2, 4},
				Cost:  6,
			},
		},
	}

	for _, c := range testcases {
		s.T().Run(c.name, func(t *testing.T) {
			g := optimization.Optimization{}
			s.Assert().Equal(
				c.expected,
				g.TravelingSalesman(c.n, c.distances, c.startAt),
			)
		})
	}
}

// var benchmarkTravelingSalesmanResult [][]int

// func BenchmarkTravelingSalesman(b *testing.B) {
// 	var r [][]int
// 	g := optimization.Optimization{}
// 	for n := 0; n < b.N; n++ {
// 		// always record the result of Fib to prevent
// 		// the compiler eliminating the function call.
// 		r = g.TravelingSalesman(30)
// 	}
// 	// always store the result to a package level variable
// 	// so the compiler cannot eliminate the Benchmark itself.
// 	benchmarkTravelingSalesmanResult = r
// }
