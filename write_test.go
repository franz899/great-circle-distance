package main

import (
	"os"
	"testing"
)

func TestWriteResultsToFile(t *testing.T) {
	name := "file.test"
	customers := []Customer{
		{1, "New York", 40.6976701, -74.2598749},
		{2, "Philadelphia", 40.0026767, -75.2581147},
		{3, "London", 51.5287352, -0.381763},
		{4, "Liverpool", 53.4123001, -3.0561377},
	}

	WriteResultsToFile(name, customers)

	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		t.Error("file should exist by now")
	}

	data, err := os.ReadFile(name)
	if err != nil {
		t.Error(err)
	}
	want := "1 New York\n2 Philadelphia\n3 London\n4 Liverpool\n"
	if string(data) != want {
		t.Errorf("WriteResultsToFile() = %v;\n want\n%v", string(data), want)
	}

	os.Remove(name)
}
