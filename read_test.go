package main

import (
	"testing"
)

func TestReadCustomers(t *testing.T) {
	want := []Customer{
		{1, "New York", 40.6976701, -74.2598749},
		{2, "Philadelphia", 40.0026767, -75.2581147},
		{3, "London", 51.5287352, -0.381763},
		{4, "Liverpool", 53.4123001, -3.0561377},
	}
	got := ReadCustomers("customers_test.txt")

	if len(got) != len(want) {
		t.Errorf(`ReadCustomers("customers_test.txt") = %v; want %v`, got, want)
	}

	for i, e := range got {
		if e != want[i] {
			t.Errorf(`ReadCustomers("customers_test.txt") = %v; want %v`, got, want)
		}
	}

}
