package optimization

type CityDistance struct {
	X        uint
	Y        uint
	Distance uint
}

type TravelRoute struct {
	Route []uint
	Cost  uint
}

func (o Optimization) TravelingSalesman(
	n uint,
	distances []CityDistance,
	startAt uint,
) TravelRoute {
	if n == 0 {
		return TravelRoute{}
	}

	var (
		route       TravelRoute   = TravelRoute{[]uint{}, ^uint(0)}
		traveledMap map[uint]bool = make(map[uint]bool)
		try         func(
			accRoute TravelRoute,
			currentCity uint,
		)
		checkCompletedAndGetPossibleNexts func(currentCity uint) (
			distanceToStartPoint uint,
			nexts map[uint]uint,
		)
	)

	try = func(accRoute TravelRoute, currentCity uint) {
		distanceToStart, nexts := checkCompletedAndGetPossibleNexts(currentCity)
		if distanceToStart > 0 &&
			len(nexts) == 0 &&
			accRoute.Cost+distanceToStart < route.Cost {
			route = TravelRoute{
				Route: append(accRoute.Route, startAt),
				Cost:  accRoute.Cost + distanceToStart,
			}
		} else {
			for city, weight := range nexts {
				if accRoute.Cost+weight > route.Cost {
					continue
				}

				traveledMap[city] = true
				try(
					TravelRoute{
						Route: append(accRoute.Route, city),
						Cost:  accRoute.Cost + weight,
					},
					city,
				)
				traveledMap[city] = false
			}
		}
	}

	checkCompletedAndGetPossibleNexts = func(currentCity uint) (
		distanceToStartPoint uint,
		nexts map[uint]uint,
	) {
		nexts = make(map[uint]uint)
		for _, d := range distances {
			switch {
			case currentCity == d.X:
				if !traveledMap[d.Y] {
					nexts[d.Y] = d.Distance
				}

				if d.Y == startAt {
					distanceToStartPoint = d.Distance
				}
			case currentCity == d.Y:
				if !traveledMap[d.X] {
					nexts[d.X] = d.Distance
				}

				if d.X == startAt {
					distanceToStartPoint = d.Distance
				}
			}
		}

		return
	}

	traveledMap[startAt] = true
	try(
		TravelRoute{append([]uint{}, startAt), 0},
		startAt,
	)

	return route
}
