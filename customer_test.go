package main

import (
	"testing"
)

func TestSortByID(t *testing.T) {
	arg := []Customer{
		{3, "London", 51.5287352, -0.381763},
		{1, "New York", 40.6976701, -74.2598749},
		{4, "Liverpool", 53.4123001, -3.0561377},
		{2, "Philadelphia", 40.0026767, -75.2581147},
	}
	want := []Customer{
		{1, "New York", 40.6976701, -74.2598749},
		{2, "Philadelphia", 40.0026767, -75.2581147},
		{3, "London", 51.5287352, -0.381763},
		{4, "Liverpool", 53.4123001, -3.0561377},
	}
	SortByID(&arg)
	if arg[0] != want[0] {
		t.Error()
	}
}

func TestFilterByDistance(t *testing.T) {
	arg := []Customer{
		{1, "Alice Cahill", 51.92893, -10.27699},
		{2, "Ian McArdle", 51.8856167, -10.4240951},
		{4, "Ian Kehoe", 53.2451022, -6.238335},
		{8, "Eoin Ahearn", 54.0894797, -6.18671},
	}
	want := []Customer{
		{4, "Ian Kehoe", 53.2451022, -6.238335},
		{8, "Eoin Ahearn", 54.0894797, -6.18671},
	}
	got := FilterByDistance(&arg)
	if len(got) != len(want) {
		t.Errorf("FilterByDistance() len = %v; want %v", len(got), len(want))
	}
}
