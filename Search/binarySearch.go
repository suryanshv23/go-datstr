package datstr

// BinarySearch is to implement binary search. It gives index and a boolean value as result.
func BinarySearch(Arr []int, item int) (int64, bool) {
	start := 0
	end := len(Arr) - 1

	for start <= end {

		mid := (start + end) / 2

		if Arr[mid] == item {
			return int64(mid), true
		} else if Arr[mid] < item {
			start = mid + 1
		} else if Arr[mid] > item {
			end = mid - 1
		}
	}
	return -1, false
}
