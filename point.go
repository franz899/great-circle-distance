package main

import "math"

const (
	EARTH_RADIUS = 6371 // in metres
)

type Point struct {
	lat float64
	lng float64
}

func CreatePoint(lat float64, lng float64) *Point {
	return &Point{lat: lat, lng: lng}
}

func (p *Point) Lat() float64 {
	return p.lat
}

func (p *Point) Lng() float64 {
	return p.lng
}

// Returns the distance in Km
func (p *Point) GetDistance(p2 *Point) float64 {
	// Formula taken from http://www.movable-type.co.uk/scripts/latlong.html
	φ1 := convertToRadians(p.Lat()) // φ, λ in radians
	φ2 := convertToRadians(p2.Lat())
	Δφ := convertToRadians(p2.Lat() - p.Lat())
	Δλ := convertToRadians(p2.Lng() - p.Lng())

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return roundToFirstDecimalPlace(EARTH_RADIUS * c)
}

func convertToRadians(c float64) float64 {
	return c * math.Pi / 180
}

func roundToFirstDecimalPlace(n float64) float64 {
	return math.Round(n*10) / 10
}
