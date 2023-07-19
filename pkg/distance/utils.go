package distance

import (
	"math"
)

const (
	radiusOfEarth = 6371 // km
)

type Point struct {
	Latitude  float64
	Longitude float64
}

func (p *Point) toRadians() {
	p.Latitude = degToRadian(p.Latitude)
	p.Longitude = degToRadian(p.Longitude)
}

func DistanceBetweenTwoPoints(p, q Point) float64 {
	p.toRadians()
	q.toRadians()

	deltaLat := p.Latitude - q.Latitude
	deltaLong := p.Longitude - q.Longitude

	a := math.Pow(
		math.Sin(deltaLat/2),
		2,
	) + math.Cos(
		p.Latitude,
	)*math.Cos(
		q.Latitude,
	)*math.Pow(
		math.Sin(deltaLong/2),
		2,
	)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return radiusOfEarth * c
}

func degToRadian(deg float64) float64 {
	return deg * math.Pi / 180
}
