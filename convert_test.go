package main

import (
	"testing"
)

func TestConvertToRadians(t *testing.T) {
	arg := 120.0
	want := 2.0943951023931953
	got := ConvertToRadians(arg)
	if got != want {
		t.Errorf("ConvertToRadians(%v) = %v; want %v", arg, got, want)
	}
}
