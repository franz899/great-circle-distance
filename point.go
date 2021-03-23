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
	φ1 := p.Lat() * math.Pi / 180 // φ, λ in radians
	φ2 := p2.Lat() * math.Pi / 180
	Δφ := (p2.Lat() - p.Lat()) * math.Pi / 180
	Δλ := (p2.Lng() - p.Lng()) * math.Pi / 180

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return roundToFirstDecimalPlace(EARTH_RADIUS * c)
}

func roundToFirstDecimalPlace(n float64) float64 {
	return math.Round(n*10) / 10
}
