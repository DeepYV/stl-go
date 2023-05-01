package search

import "errors"

func LinearSearch(arr []int, target int) (int, error) {

	for i := 0; i < len(arr); i++ {

		if target == arr[i] {
			return i, nil
		}
	}

	return -1, errors.New("element not found")
}
