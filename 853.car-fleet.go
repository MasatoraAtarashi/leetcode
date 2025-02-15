package main

import (
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	type car struct {
		position int
		speed    int
	}
	cars := make([]car, len(position))
	for i := 0; i < len(position); i++ {
		cars[i] = car{
			position: position[i],
			speed:    speed[i],
		}
	}

	// ターゲットに近い順にソート
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	fleet := Stack[float64]{}
	for _, car := range cars {
		goalTime := float64(target-car.position) / float64(car.speed)
		if !fleet.Empty() && goalTime <= fleet.Top() {
			continue
		}
		fleet.Push(goalTime)
	}

	return len(fleet)
}
