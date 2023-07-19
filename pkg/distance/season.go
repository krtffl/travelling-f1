package distance

import (
	"log"
	"strconv"

	"github.com/krtffl/travelling-f1/pkg/ergast"
)

func GetLocation(race ergast.Race) Point {
	lat, err := strconv.ParseFloat(race.Circuit.Location.Lat, 64)
	if err != nil {
		log.Fatalln(err)
	}

	long, err := strconv.ParseFloat(race.Circuit.Location.Long, 64)
	if err != nil {
		log.Fatalln(err)
	}

	return Point{
		Latitude:  lat,
		Longitude: long,
	}
}

func DistanceBetweenTwoRaces(x, y ergast.Race) float64 {
	p := GetLocation(x)
	q := GetLocation(y)

	distance := DistanceBetweenTwoPoints(p, q)
	return distance
}

func SeasonDistance(races []ergast.Race) float64 {
	var prev ergast.Race
	var total float64
	for _, r := range races {
		if prev != (ergast.Race{}) {
			total += DistanceBetweenTwoRaces(prev, r)
		}
		prev = r
	}
	return total
}
