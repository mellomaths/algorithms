package search

func BinarySearch(array []int, target int) (int, error) {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] == target {
			return mid, nil
		} else if array[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, ErrNotFound
}
