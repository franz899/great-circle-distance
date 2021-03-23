package main

func main() {
	customers := ReadCustomers("customers.txt")

	closeCustomers := FilterByDistance(&customers)

	SortByID(&closeCustomers)

	WriteResultsToFile("output.txt", closeCustomers)
}
