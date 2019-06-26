package datstr

//LinearSearch is for searching an array linearly
func LinearSearch(Arr []int, item int) (int64, bool) {
	for index, val := range Arr {
		if val == item {
			return int64(index), true
		}
	}
	return -1, false
}
