package main

import "sort"

var (
	Dublin = Point{53.339428, -6.257664}
)

type Customer struct {
	ID        int     `json:"user_id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
}

func SortByID(customers *[]Customer) {
	sort.Slice((*customers), func(i, j int) bool {
		return (*customers)[i].ID < (*customers)[j].ID
	})
}

func FilterByDistance(customers *[]Customer) []Customer {
	filtered := make([]Customer, 0)
	for _, customer := range *customers {
		p := Point{customer.Latitude, customer.Longitude}
		distance := Dublin.GetDistance(&p)

		if distance < 101 {
			filtered = append(filtered, customer)
		}
	}
	return filtered
}
