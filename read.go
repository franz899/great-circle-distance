package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"strings"
)

// Read customers' file and return a list of them
func ReadCustomers(name string) []Customer {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))

	customers := make([]Customer, 0, 50)

	for scanner.Scan() {
		c := Customer{}
		err := json.Unmarshal(scanner.Bytes(), &c)
		if err != nil {
			panic(err)
		}
		customers = append(customers, c)
	}

	return customers
}
