package main

import (
	"fmt"
	binarySearch "stl/binarySearch"
)

func main() {

	a, err := binarySearch.BinarySearch()([]int{1, 2, 3, 4, 5}, 7)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)

}
