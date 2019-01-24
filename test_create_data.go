package main

import (
	"market-server/server/model"
)

func main() {
	model.CreateTestDataFromFile("./test_data.dat")
}
