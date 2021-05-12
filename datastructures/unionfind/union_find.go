package unionfind

type UnionFind struct {
	parent []int
	rank   []int
}

// New return a pointer to UnionFind which specific size
func New(size int) *UnionFind {
	return new(UnionFind).init(size)
}

func (uf *UnionFind) init(size int) *UnionFind {
	uf = new(UnionFind)

	uf.parent = make([]int, size)
	uf.rank = make([]int, size)

	for i := 0; i < size; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}

	return uf
}

// Find using Path compression technique,
// which involves changing the x = parent[x] in the find function to parent[x] = find(parent[x]).
// Basically, as we compute the correct leader for x, we should remember our calculation.
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	return uf.parent[x]
}

// Union use Union-by-rank technique,
// which involves distributing the workload of find across leaders evenly.
// Whenever we call Union(x, y), we have two leaders xr, yr
// and we have to choose whether we want parent[x] = yr or parent[y] = xr.
// We choose the leader that has a higher following to pick up a new follower.
// Specifically, the meaning of rank is that there are less than 2 ^ rank[x] followers of x.
// This strategy can be shown to give us better bounds for how long the recursive loop in Find could run for.
func (uf *UnionFind) Union(x, y int) bool {
	xr, yr := uf.Find(x), uf.Find(y)
	if xr == yr {
		return false
	}

	if uf.rank[xr] < uf.rank[yr] {
		uf.parent[xr] = yr
	} else if uf.rank[xr] > uf.rank[yr] {
		uf.parent[yr] = xr
	} else {
		uf.parent[yr] = xr
		uf.rank[xr] += 1
	}

	return true
}
