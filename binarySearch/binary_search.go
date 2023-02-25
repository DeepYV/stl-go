package bs

import "errors"

type Binary func(arr []int, target int) (int, error)

func BinarySearch() Binary {
	return Binary(func(arr []int, target int) (int, error) {
		left := 0
		right := len(arr) - 1

		for left <= right {
			mid := (left + right) / 2
			if arr[mid] == target {
				return mid, nil
			} else if arr[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return -1, errors.New("element not found")

	})
}
