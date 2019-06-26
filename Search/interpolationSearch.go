package datstr

// InterpolationSearch is for searchin an Array using interplation formulae.
//It gives index and a boolean value as result.
func InterpolationSearch(Arr []int, item int) (int64, bool) {
	start, end := 0, len(Arr)-1

	for start <= end {

		if start == end {
			if item == Arr[start] {
				return int64(start), true
			}
			return -1, false
		}

		pos := start + (((end - start) / (Arr[end] - Arr[start])) * (item - Arr[start]))

		if pos < 0 {
			return -1, false
		}

		if Arr[pos] == item {
			return int64(pos), true
		} else if Arr[pos] < item {
			start = pos + 1
		} else if Arr[pos] > item {
			end = pos - 1
		}
	}

	return -1, false
}
