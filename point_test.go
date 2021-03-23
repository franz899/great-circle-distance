package main

import "testing"

func TestCreatePoint(t *testing.T) {
	want := &Point{10.0, 11.0}
	got := CreatePoint(10.0, 11.0)
	if got.lat != want.lat || got.lng != want.lng {
		t.Errorf("CreatePoint(10.0, 11.0) = %v; want %v", got, want)
	}
}

func TestLat(t *testing.T) {
	p := &Point{10.0, 11.0}
	want := 10.0
	got := p.Lat()
	if got != want {
		t.Errorf("Lat() = %v; want %v", got, want)
	}
}

func TestLng(t *testing.T) {
	p := &Point{10.0, 11.0}
	want := 11.0
	got := p.Lng()
	if got != want {
		t.Errorf("Lng() = %v; want %v", got, want)
	}
}

func TestGetDistance(t *testing.T) {
	newYork := &Point{40.6976701, -74.2598749}
	philadelphia := &Point{40.0026767, -75.2581147}
	want := 114.6
	got := newYork.GetDistance(philadelphia)
	if got != want {
		t.Errorf("GetDistance(philadelphia) = %v; want %v", got, want)
	}

	london := &Point{51.5287352, -0.381763}
	liverpool := &Point{53.4123001, -3.0561377}
	want = 276.9
	got = london.GetDistance(liverpool)
	if got != want {
		t.Errorf("GetDistance(liverpool) = %v; want %v", got, want)
	}
}
