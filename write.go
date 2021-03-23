package main

import (
	"bufio"
	"fmt"
	"os"
)

func WriteResultsToFile(name string, customers []Customer) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, customer := range customers {
		_, err := w.WriteString(fmt.Sprintf("%v %v\n", customer.ID, customer.Name))
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
}
