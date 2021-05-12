package dp

// There are persons and tasks, each task is to be alloted to a single person.
// We are also given a matrix  of size , where  denotes, how much person  is going to charge for task .
// Now we need to assign each task to a person in such a way that the total cost is minimum.
// Note that each task is to be alloted to a single person, and each person will be alloted only one task.
type AssignTask struct {
}

func (a AssignTask) BruteForce(n int, costs [][]int) int {
	minCost := maxInt
	assignedTaskToPersonMap := make(map[int]int)
	var assign func(person int, cost int)

	min := func(a, b int) int {
		if a >= b {
			return a
		}

		return b
	}

	assign = func(person int, cost int) {
		// base case
		if person == n+1 {
			minCost = min(minCost, cost)
		} else {
			for j := 1; j <= n; j++ {
				if _, ok := assignedTaskToPersonMap[j]; ok {
					continue
				}

				assignedTaskToPersonMap[j] = person
				assign(person+1, cost+costs[person-1][j-1])
				delete(assignedTaskToPersonMap, j)
			}
		}
	}

	assign(1, 0)

	return minCost
}
