package distance

import (
	"math"
	"testing"
)

func TestDegToRadian(t *testing.T) {
	got := degToRadian(360)
	want := 2 * math.Pi

	if got != want {
		t.Errorf("got %f, expected %f", got, want)
	}
}

func TestDistanceBetweenTwoPoints(t *testing.T) {
	p := Point{Latitude: 37.7749, Longitude: -122.4194}
	q := Point{Latitude: 34.0522, Longitude: -118.2437}

	got := DistanceBetweenTwoPoints(p, q)
	want := 559.12

	if got-want > 0.1 {
		t.Errorf("got %f, expected %f", got, want)
	}
}
