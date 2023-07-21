package distance

import (
	"testing"

	"github.com/krtffl/travelling-f1/pkg/ergast"
)

func race(lat, long string) ergast.Race {
	return ergast.Race{
		Circuit: ergast.Circuit{
			Location: ergast.Location{
				Lat:  lat,
				Long: long,
			},
		},
	}
}

func TestGetLocation(t *testing.T) {
	r := race("69.69", "420.69")
	got := GetLocation(r)

	want := Point{
		Latitude:  69.69,
		Longitude: 420.69,
	}

	if got != want {
		t.Errorf("got %f, expected %f", got, want)
	}
}

func TestDistanceBetweenTwoRaces(t *testing.T) {
	firstRace := race("37.7749", "-122.4194")
	secondRace := race("34.0522", "-118.2437")

	got := DistanceBetweenTwoRaces(firstRace, secondRace)
	want := 559.12

	if got-want > 0.1 {
		t.Errorf("got %f, expected %f", got, want)
	}
}

func TestSeasonDistance(t *testing.T) {
	firstRace := race("37.7749", "-122.4194")
	secondRace := race("34.0522", "-118.2437")
	thirdRace := race("37.7749", "-122.4194")

	season := []ergast.Race{firstRace, secondRace, thirdRace}

	got := SeasonDistance(season)
	want := 559.12 * 2

	if got-want > 0.1 {
		t.Errorf("got %f, expected %f", got, want)
	}
}
