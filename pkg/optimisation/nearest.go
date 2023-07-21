package optimisation

import (
	"math"

	"github.com/krtffl/travelling-f1/pkg/distance"
	"github.com/krtffl/travelling-f1/pkg/ergast"
)

func NearestNeighbour(season []ergast.Race, openingRace int) []ergast.Race {
	currentRace := season[openingRace]
	remainingRaces := append(season[:openingRace], season[openingRace+1:]...)

	path := make([]ergast.Race, 0, len(season))
	path = append(path, currentRace)

	for len(remainingRaces) > 0 {
		race, idx := findNearestRace(currentRace, remainingRaces)
		remainingRaces = append(remainingRaces[:idx], remainingRaces[idx+1:]...)
		currentRace = race
		path = append(path, currentRace)

	}
	return path
}

func findNearestRace(current ergast.Race, remaining []ergast.Race) (ergast.Race, int) {
	minDistance := math.MaxFloat64
	minIndex := 0

	for i, r := range remaining {
		distance := distance.DistanceBetweenTwoRaces(current, r)
		if distance < minDistance {
			minDistance = distance
			minIndex = i
		}
	}

	return remaining[minIndex], minIndex
}
