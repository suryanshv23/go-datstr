package datstr

// InsertionSort is for sortig array
func InsertionSort(Arr []int) []int {
	n := len(Arr)
	for i := 0; i <= n-2; i++ {
		j := i
		for j >= 0 {
			if Arr[j] > Arr[j+1] {
				Arr[j], Arr[j+1] = Arr[j+1], Arr[j]
			}
			j = j - 1
		}
	}
	return Arr
}
