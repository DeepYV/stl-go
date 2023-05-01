package main

import (
	"fmt"

	"github.com/DeepYV/stl-go/search"
)

func main() {

	// binary search
	a, err := search.BinarySearch()([]int{1, 2, 3, 4, 5}, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)

	// linear search
	b, err := search.LinearSearch([]int{1, 2, 3, 4, 5}, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(b)

}
