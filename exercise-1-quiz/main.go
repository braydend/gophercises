package main

import (
	"fmt"
)

func main() {
	fmt.Println(readCsv(getCsvReader("addition.csv")))
}
