package main

import "math"

// Convert GPS coordinate to radians
func ConvertToRadians(c float64) float64 {
	return c * math.Pi / 180
}
